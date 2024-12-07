package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const URL = "http://localhost:4000/greet"

func TestReponseOK(t *testing.T) {

	// arrange
	query := "?name=Test"
	req := httptest.NewRequest(http.MethodGet, URL+query, nil)
	rr := httptest.NewRecorder()

	validateRequestHandler := func(_ http.ResponseWriter, req *http.Request) {

		// check on correct request method
		if req.Method != http.MethodGet {
			http.Error(rr, "Wrong method used! Only GET method allowed.", http.StatusMethodNotAllowed)
			return
		}

		// parse encoded query params to map
		params, err := url.ParseQuery(req.URL.Query().Encode())
		if err != nil {
			http.Error(rr, "Invalid query parameters!", http.StatusBadRequest)
			return
		}

		// check if "name" word is exists
		if _, ok := params[string(contextNameKey)]; !ok {
			http.Error(rr, "'name' parameter does not exist or bad value!", http.StatusBadRequest)
			return
		}

		// parse search query
		text := params.Get(string(contextNameKey))
		if text == "" {
			http.Error(rr, "'name' value is empty!", http.StatusBadRequest)
			return
		}
	}

	// act
	handler := ValidateRequest(http.HandlerFunc(validateRequestHandler))
	handler.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	// assert
	if resp.StatusCode != http.StatusOK || req.Method != http.MethodGet {
		t.Errorf("Bad response status code! Excpect: %d Have: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestReponseBadRequest(t *testing.T) {

	// arrange
	query := "?name=" //TODO: add more cases
	req := httptest.NewRequest(http.MethodGet, URL+query, nil)
	rr := httptest.NewRecorder()

	validateRequestHandler := func(_ http.ResponseWriter, req *http.Request) {

		// parse encoded query params to map
		params, _ := url.ParseQuery(req.URL.Query().Encode())

		// check if "name" word is exists
		if _, ok := params[string(contextNameKey)]; !ok {
			http.Error(rr, "'name' parameter does not exist or bad value!", http.StatusBadRequest)
			return
		}
	}

	// act
	handler := ValidateRequest(http.HandlerFunc(validateRequestHandler))
	handler.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	// assert
	if resp.StatusCode != http.StatusBadRequest || req.Method != http.MethodGet {
		t.Errorf("Bad response status code! Excpect: %d Have: %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestWrongMethodValidStatusCode(t *testing.T) {

	// arrange
	query := "?name=Test"
	req := httptest.NewRequest(http.MethodPost, URL+query, nil)
	rr := httptest.NewRecorder()

	validateRequestHandler := func(_ http.ResponseWriter, req *http.Request) {

		// check on correct request method
		if req.Method != http.MethodGet {
			http.Error(rr, "Wrong method used! Only GET method allowed.", http.StatusMethodNotAllowed)
			return
		}
	}

	// act
	handler := ValidateRequest(http.HandlerFunc(validateRequestHandler))
	handler.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	// assert
	if resp.StatusCode != http.StatusMethodNotAllowed || req.Method != http.MethodPost {
		t.Errorf("Bad response status code! Excpect: %d Have: %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestInvalidQuery(t *testing.T) {

	// arrange
	query := "?n"
	req := httptest.NewRequest(http.MethodGet, URL+query, nil)
	rr := httptest.NewRecorder()

	validateRequestHandler := func(_ http.ResponseWriter, req *http.Request) {

		// parse encoded query params to map
		_, err := url.ParseQuery(req.URL.Query().Encode())
		if err != nil {
			http.Error(rr, "Invalid query parameters!", http.StatusBadRequest)
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
