package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/logging"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
)

func main() {
	prj := os.Getenv("PROJECT")
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: prj,
	})
	if err != nil {
		panic(err)
	}
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	http.HandleFunc("/", index())
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}

func index() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hf := &propagation.HTTPFormat{}
		sc, ok := hf.SpanContextFromRequest(r)
		if !ok {
			fmt.Fprintf(w, "error propagation")
		}
		ctx, _ := trace.StartSpanWithRemoteParent(r.Context(), "index", sc)

		lc, err := logging.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
		if err != nil {
			log.Println(err)
			return
		}
		defer lc.Close()

		opts := logging.CommonResource(
			&mrpb.MonitoredResource{
				Type: "gae_app",
				Labels: map[string]string{
					"project_id": os.Getenv("GOOGLE_CLOUD_PROJECT"),
					"module_id":  os.Getenv("GAE_SERVICE"),
					"version_id": os.Getenv("GAE_VERSION"),
					"zone":       "asia-northeast1-1",
				},
			},
		)
		lg := lc.Logger("mysample", opts)
		traceId := [16]byte(sc.TraceID)
		js := []byte(`{"Name": "test", "Count": 2}`)
		lg.Log(logging.Entry{Payload: json.RawMessage(js), Severity: logging.Info, Trace: "projects/dena-internal-gdi-gcp/traces/" + hex.EncodeToString(traceId[:])})
		lg.Log(logging.Entry{Payload: "hogehoge", Severity: logging.Warning, Trace: "projects/dena-internal-gdi-gcp/traces/" + hex.EncodeToString(traceId[:])})
		lg.StandardLogger(logging.Info).Println("hello")
		lg.StandardLogger(logging.Warning).Println("hello")
		log.Println("normal log")
		fmt.Println("normal log")
		longProcess(lg)
		fmt.Fprintf(w, "hello 2nd-gen")
	}
}

func longProcess(lg *logging.Logger) {
	for i := 0; i < 10; i++ {
		lg.StandardLogger(logging.Info).Println("count", i)
		time.Sleep(time.Second * 1)
	}
}
