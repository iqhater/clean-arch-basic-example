package controller

import (
	infra "clean_arch_basic_example/internal/infrastructure"
	"clean_arch_basic_example/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGreetPrivateHandlerValid(t *testing.T) {

	// arrange
	greeterPrivateRepoMock := &infra.GreetPrivateMockDB{
		IDs: map[int]infra.GreetingsCounter{
			42: 0,
		},
	}
	greeterPrivateServiceMock := usecase.NewPrivateGreetService(greeterPrivateRepoMock)
	greetPrivateController := NewPrivateGreetController(greeterPrivateServiceMock)

	name := "Bob"
	userID := "42"

	id, err := strconv.Atoi(userID)
	if err != nil {
		t.Errorf("'id' value must be integer! %v", err)
	}

	ctx := context.WithValue(context.Background(), contextNameKey, name)
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/greet:id", nil)
	req.SetPathValue(string(contextIDKey), userID)

	rr := httptest.NewRecorder()

	// act
	greetPrivateController.GreetPrivateHandler(rr, req)

	// assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returns wrong status code: got %v want %v", status, http.StatusOK)
	}

	output := ResponsePrivateGreetDTO{
		ID:      id,
		Title:   infra.GREET_PRIVATE_TITLE + " " + name + "!",
		Counter: 1,
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		t.Errorf("Cannot parse json struct! %v", err)
	}

	expected := string(outputJson)
	if rr.Body.String() != expected {
		t.Errorf("handler returns unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

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
