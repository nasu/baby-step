package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
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
		sc, err := getSpanContext(r)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		ctx, _ := trace.StartSpanWithRemoteParent(r.Context(), "index", *sc)
		lc, err := getLoggerClient(ctx)
		if err != nil {
			fmt.Fprint(w, err)
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
		//variety(lg, getTraceID(*sc))
		longString(lg)
		//longProcess(lg)
		fmt.Fprintf(w, "hello 2nd-gen")
	}
}

func getSpanContext(r *http.Request) (*trace.SpanContext, error) {
	hf := &propagation.HTTPFormat{}
	sc, ok := hf.SpanContextFromRequest(r)
	if !ok {
		return nil, errors.New("Failed to propagation")
	}
	return &sc, nil
}

func getLoggerClient(ctx context.Context) (*logging.Client, error) {
	lc, err := logging.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, err
	}
	return lc, nil
}

func getTraceID(sc trace.SpanContext) string {
	traceID := [16]byte(sc.TraceID)
	return hex.EncodeToString(traceID[:])
}

func variety(lg *logging.Logger, traceID string) {
	js := []byte(`{"Name": "test", "Count": 2}`)
	lg.Log(logging.Entry{Payload: json.RawMessage(js), Severity: logging.Info, Trace: "projects/" + os.Getenv("GOOGLE_CLOUD_PROJECT") + "/traces/" + traceID})
	lg.Log(logging.Entry{Payload: "hogehoge", Severity: logging.Warning, Trace: "projects/" + os.Getenv("GOOGLE_CLOUD_PROJECT") + "/traces/" + traceID})
	lg.StandardLogger(logging.Info).Println("hello")
	lg.StandardLogger(logging.Warning).Println("hello")
	log.Println("normal log")
	fmt.Println("normal log")
}

func longProcess(lg *logging.Logger) {
	for i := 0; i < 10; i++ {
		lg.StandardLogger(logging.Info).Println("count", i)
		time.Sleep(time.Second * 1)
	}
}

func longString(lg *logging.Logger) {
	var str string
	for i := 0; i < 1000*100; i++ {
		str += "0"
	}
	lg.Log(logging.Entry{Payload: str, Severity: logging.Info})
	lg.Log(logging.Entry{Payload: fmt.Sprintf("len:%d", len(str)), Severity: logging.Info})
	/*
	   if err := lg.Flush(); err != nil {
	       log.Println("ERROR:", err)
	   }
	*/
}
