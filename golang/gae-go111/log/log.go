package log

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func NewWithSugar(sugar *zap.SugaredLogger) *Logger {
	return &Logger{sugar}
}

// Debugf is to output a message with DEBUG severity
func (l Logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.sugar.Debugw(msg, "foo", "bar")
}

// Infof is to output a message with INFO severity
func (l Logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.sugar.Infow(msg, "foo", "bar")
}

// Warnf is to output a message with WARN severity
func (l Logger) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.sugar.Warnw(msg, "foo", "bar")
}

// Errorf is to output a message with ERROR severity
func (l Logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.sugar.Errorw(msg, "foo", "bar")
}

// Criticalf is to output a message with CRITICAL severity
func (l Logger) Criticalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.sugar.DPanicw(msg, "foo", "bar")
}
