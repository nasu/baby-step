package log

import (
	zapstackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

func GetInstance() (*Logger, error) {
	if logger == nil {
		if err := instantiate(); err != nil {
			return nil, err
		}
	}
	return logger, nil
}

func ResetInstance() {
	wrapInstance(nil)
}

func wrapInstance(l *Logger) {
	logger = l
}

func instantiate() error {
	conf := &zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		EncoderConfig:    zapstackdriver.EncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zaplogger, err := conf.Build(
		zap.WrapCore(
			func(core zapcore.Core) zapcore.Core {
				return &zapstackdriver.Core{Core: core}
			},
		),
		zap.Fields(
			zapstackdriver.LogServiceContext(&zapstackdriver.ServiceContext{
				Service: "foo",
				Version: "bar",
			}),
		),
	)
	if err != nil {
		return err
	}
	logger = &Logger{
		sugar: zaplogger.Sugar(),
	}
	return nil
}
