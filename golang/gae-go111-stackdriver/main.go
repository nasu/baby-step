package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
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
		hf := &propagation.HTTPFormat{}
		sc, ok := hf.SpanContextFromRequest(r)
		if !ok {
			fmt.Fprintf(w, "error propagation")
		}
		ctx, span := trace.StartSpanWithRemoteParent(r.Context(), "index", sc)
		defer span.End()
		process1(ctx)
		process2(ctx)
		longProcess(ctx)
		fmt.Fprintf(w, "hello 2nd-gen")
	}
}

func process1(ctx context.Context) {
	_, span := trace.StartSpan(ctx, "index.process1")
	span.Annotatef(nil, "start process1")
	time.Sleep(time.Second * 1)
	span.Annotatef(nil, "end process1")
	span.End()
}

func process2(ctx context.Context) {
	_, span := trace.StartSpan(ctx, "index.process2")
	span.Annotatef(nil, "start process2")
	time.Sleep(time.Second * 1)
	span.Annotatef(nil, "end process2")
	span.End()
}

func longProcess(ctx context.Context) {
	_, span := trace.StartSpan(ctx, "index.longProcess")
	for i := 0; i < 10; i++ {
		attrs := []trace.Attribute{
			trace.StringAttribute("string", "foobar"),
			trace.BoolAttribute("bool", true),
			trace.Int64Attribute("int", int64(12345)),
		}
		span.Annotatef(attrs, fmt.Sprintf("count: %d", i))
		time.Sleep(time.Second * 1)
	}
	span.End()
}
