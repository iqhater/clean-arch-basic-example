package status

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeaderValid(t *testing.T) {

	sr := StatusHTTP{
		httptest.NewRecorder(),
		http.StatusOK,
	}

	testStatusCode := 400
	sr.WriteHeader(testStatusCode)

	if sr.StatusCode != testStatusCode {
		t.Errorf("Wrong Status Code in header!: got %d", sr.StatusCode)
	}
}

func TestNotEmptyNewStatusHTTP(t *testing.T) {
	rr := httptest.NewRecorder()
	result := NewStatusHTTP(rr)

	if result == nil {
		t.Errorf("NewStatusHTTP must return non nil!: got %v", result)
	}
}
