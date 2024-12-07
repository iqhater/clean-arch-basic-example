package controller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// TODO: mock GreetPrivateHandler
func TestLimitIDNotNumber(t *testing.T) {

	// arrange
	query := "/iq?name=Test"
	req := httptest.NewRequest(http.MethodGet, URL+query, nil)
	rr := httptest.NewRecorder()

	validateRequestHandler := func(_ http.ResponseWriter, req *http.Request) {

		// convert id value to string
		_, err := strconv.Atoi(req.PathValue(string(contextIDKey)))
		if err != nil {
			http.Error(rr, "'id' value must be integer!", http.StatusBadRequest)
			return
		}
	}

	// act
	handler := ValidateRequest(http.HandlerFunc(validateRequestHandler))
	handler.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	// assert
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Bad response status code! Excpect: %d Have: %d", http.StatusBadRequest, resp.StatusCode)
	}
}
