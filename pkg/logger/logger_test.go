package logger

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestLogMiddleware(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	logHandler := func(_ http.ResponseWriter, req *http.Request) {
		tn := time.Now()
		log.Printf("%s %s - %s - %d - %v\n", req.Method, req.URL.String(), req.Host, http.StatusBadRequest, time.Since(tn))
	}

	rr := httptest.NewRecorder()

	handler := Log(http.HandlerFunc(logHandler))
	handler.ServeHTTP(rr, req)

	if buf.Len() == 0 {
		t.Error("Empty log output!")
	}
}
