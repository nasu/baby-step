package log_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	log "github.com/nasu/baby-step/golang/gae-go111/log"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		test func(l *log.Logger)
		want string
	}{
		{
			name: "Severity: DEBUG",
			test: func(l *log.Logger) { l.Debugf("This is a debug message.") },
			want: "",
		},
		/*
		   {
		       name: "Severity: INFO",
		       test: func(l *log.Logger) { l.Infof("This is a info message.") },
		       want: "",
		   },
		   {
		       name: "Severity: WARNING",
		       test: func(l *log.Logger) { l.Warnf("This is a info message.") },
		       want: "",
		   },
		   {
		       name: "Severity: ERROR",
		       test: func(l *log.Logger) { l.Errorf("This is a info message.") },
		       want: "",
		   },
		   {
		       name: "Severity: CRITICAL",
		       test: func(l *log.Logger) { l.Criticalf("This is a info message.") },
		       want: "",
		   },
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := captureOutput(func() error {
				logger, err := log.GetInstance()
				if err != nil {
					return err
				}
				tt.test(logger)
				return nil
			})
			if err != nil {
				t.Error(err)
				return
			}
			if w, g := tt.want, out; w != g {
				t.Errorf("wrong message. want=%s, got=%s", w, g)
			}
			log.ResetInstance()
		})
	}
}

func captureOutput(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err := f(); err != nil {
		return "", err
	}

	ch := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		ch <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-ch
	return out, nil
}
