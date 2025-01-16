package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	infra "clean_arch_basic_example/internal/infrastructure"
	srv "clean_arch_basic_example/internal/usecase"
	mid "clean_arch_basic_example/pkg/middleware"

	"github.com/google/uuid"
)

func TestGreetHandlerValid(t *testing.T) {

	// arrange
	greeterRepoMock := &infra.GreetMockDB{}
	greeterServiceMock := srv.NewGreetService(greeterRepoMock)
	greetController := NewGreetController(greeterServiceMock)

	name := "Bob"
	requestID := uuid.New()

	ctx := context.WithValue(context.Background(), mid.RequestIDKey, requestID)
	ctx2 := context.WithValue(ctx, contextNameKey, name)
	req := httptest.NewRequestWithContext(ctx2, http.MethodGet, "/greet", nil)
	rr := httptest.NewRecorder()

	// act
	greetController.GreetHandler(rr, req)

	// assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returns wrong status code: got %v want %v", status, http.StatusOK)
	}

	output := ResponsePublicGreetDTO{
		RequestID: requestID,
		Title:     infra.GREET_TITLE + " " + name + "!",
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
