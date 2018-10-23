package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func main() {
	prj := os.Getenv("PROJECT")
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: prj,
		//BundleDelayThreshold: time.Second / 10,
		//BundleCountThreshold: 10,
	})
	if err != nil {
		panic(err)
	}
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	http.HandleFunc("/", index(prj))
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}

func index(prj string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(context.Background(), "mysample")
		span.Annotatef(nil, "start index page")
		longProcess(ctx)
		fmt.Fprintf(w, "hello 2nd-gen")
		span.Annotatef(nil, "end index page")
		span.End()
	}
}

func longProcess(ctx context.Context) {
	span := trace.FromContext(ctx)
	for i := 0; i < 10; i++ {
		attrs := []trace.Attribute{
			trace.StringAttribute("string", "foobar"),
			trace.BoolAttribute("bool", true),
			trace.Int64Attribute("int", int64(12345)),
		}
		span.Annotatef(attrs, fmt.Sprintf("count: %d", i))
		time.Sleep(time.Second * 1)
	}
}
