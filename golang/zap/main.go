package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	example()
	development()
	production()
	customBasic()
}

func contextWithSugar(sugar *zap.SugaredLogger, mode string) context.Context {
	ctx := context.WithValue(context.Background(), "sugar", sugar)
	ctx = context.WithValue(ctx, "mode", mode)
	return ctx
}

func example() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	trace1(contextWithSugar(sugar, "example"))
}

func development() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	defer sugar.Sync()
	trace1(contextWithSugar(sugar, "dev"))
}

func production() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer sugar.Sync()
	trace1(contextWithSugar(sugar, "prod"))
}

func customBasic() {
	zap.NewAtomicLevel()
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "/tmp/zap.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			// Not output if not set.
			// In other words, required MessageKey at least.
			TimeKey:    "TimeStamp",
			LevelKey:   "Level",
			CallerKey:  "Caller",
			MessageKey: "Message",
			// Required. Panic if not set.
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := conf.Build()
	sugar := logger.Sugar()
	defer sugar.Sync()
	trace1(contextWithSugar(sugar, "custom"))
}

func trace1(ctx context.Context) {
	trace2(ctx)
}

func trace2(ctx context.Context) {
	trace3(ctx)
}

func trace3(ctx context.Context) {
	sugar := ctx.Value("sugar").(*zap.SugaredLogger)
	mode := ctx.Value("mode").(string)

	log.Printf("std.Log: %s", time.Now().Local())
	sugar.Debugw("Debug", "time", time.Now().Local(), "mode", mode)
	sugar.Infow("Info", "time", time.Now().Local(), "mode", mode)

	// If dev or prod mode, output a stacktrace.
	// e.g.
	// --
	// 2018-10-18T18:46:30.866+0900    WARN    zap/main.go:79  Warn    {"time": "2018-10-18T18:46:30.866+0900", "mode": "dev"}
	// main.trace3
	// baby-step/golang/zap/main.go:79
	// main.trace2
	// 		baby-step/golang/zap/main.go:67
	// --
	sugar.Warnw("Warn", "time", time.Now().Local(), "mode", mode)

	// If prod mode, output a stacktrace.
	// When prod mode, a stacktrace is into one line in other words include json
	// and a timestamp and a msg etc... are include json.
	// e.g.
	// {"level":"error","ts":1539855990.870307,"caller":"zap/main.go:83","msg":"Error","time":1539855990.870305,"mode":"prod","stacktrace":"main.trace3\n\t/baby-step/golang/zap/main.go:83\nmain.trace2\n\t/main.go:67\nmain.trace1\n\t/baby-step/golang/zap/main.go:63\nmain.production\n\t/baby-step/golang/zap/main.go:42\nmain.main\n\t/baby-step/golang/zap/main.go:15\nruntime.main\n\t/ions/1.11.1/src/runtime/proc.go:201"}
	sugar.Errorw("Error", "time", time.Now().Local(), "mode", mode)

	// Means dev panic. Panic only when dev mode.
	//sugar.DPanicw("DPanic", "time", time.Now().Local(), "mode", mode)
	//sugar.Panicw("Panic", "time", time.Now().Local(), "mode", mode)
	//sugar.Fatalw("Fatal", "time", time.Now().Local(), "mode", mode)
	//time.Sleep(time.Second * 5)
}
