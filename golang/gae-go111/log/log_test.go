package log_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	log "github.com/nasu/baby-step/golang/gae-go111/log"
)

// zap に依存したテスト
// zap 以外を採用したときはこのテストを破棄すること
func TestPrintersWithZapObserver(t *testing.T) {
	tests := []struct {
		name        string
		severity    zapcore.Level
		test        func(l *log.Logger)
		wantLen     int
		wantMessage string
	}{
		{
			name:        "Severity: DEBUG Output: DEBUG",
			severity:    zapcore.DebugLevel,
			test:        func(l *log.Logger) { l.Debugf("This is a debug message.") },
			wantLen:     1,
			wantMessage: "This is a debug message.",
		},
		{
			name:        "Severity: INFO Output: INFO",
			severity:    zapcore.InfoLevel,
			test:        func(l *log.Logger) { l.Infof("This is a info message.") },
			wantLen:     1,
			wantMessage: "This is a info message.",
		},
		{
			name:        "Severity: INFO Output: WARN",
			severity:    zapcore.InfoLevel,
			test:        func(l *log.Logger) { l.Warnf("This is a warn message.") },
			wantLen:     1,
			wantMessage: "This is a warn message.",
		},
		{
			name:        "Severity: INFO Output: ERROR",
			severity:    zapcore.InfoLevel,
			test:        func(l *log.Logger) { l.Errorf("This is a error message.") },
			wantLen:     1,
			wantMessage: "This is a error message.",
		},
		{
			name:        "Severity: INFO Output: CRITICAL",
			severity:    zapcore.InfoLevel,
			test:        func(l *log.Logger) { l.Criticalf("This is a critical message.") },
			wantLen:     1,
			wantMessage: "This is a critical message.",
		},
		{
			name:        "Severity: INFO Output: DEBUG",
			severity:    zapcore.InfoLevel,
			test:        func(l *log.Logger) { l.Debugf("This is a debug message.") },
			wantLen:     0,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, observed := observer.New(tt.severity)
			zaplogger := zap.New(core)
			tt.test(log.NewWithSugar(zaplogger.Sugar()))
			if w, g := tt.wantLen, observed.Len(); w != g {
				t.Errorf("wrong length. want=%d, got=%d", w, g)
			}
			if 0 < tt.wantLen {
				if w, g := tt.wantMessage, observed.AllUntimed()[0].Message; w != g {
					t.Errorf("wrong message. want=%s, got=%s", w, g)
				}
			}
		})
	}
}
