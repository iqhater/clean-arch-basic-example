package controller

import (
	srv "clean_arch_basic_example/internal/usecase"
	"clean_arch_basic_example/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const contextNameKey contextKey = "name"

type ResponsePublicGreetDTO struct {
	RequestID uuid.UUID `json:"request_id"`
	Title     string    `json:"greeting"`
}

// controller/handler
type GreetController struct {
	greetService *srv.GreetService
}

// init
func NewGreetController(greetService *srv.GreetService) *GreetController {
	return &GreetController{
		greetService: greetService,
	}
}

// implementation
func (g *GreetController) GreetHandler(w http.ResponseWriter, req *http.Request) {

	// add unique request id from context
	requestID, ok := req.Context().Value(pkg.RequestIDKey).(uuid.UUID)
	if !ok || requestID == uuid.Nil {
		fmt.Fprintln(w, "this a request without requestID")
		return
	}

	// name and method request validations are located in separate validateRequest middleware
	// get id and name params from context
	name, ok := req.Context().Value(contextNameKey).(string)
	if !ok {
		fmt.Fprintln(w, "this a request without name")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	result, err := g.greetService.GetGreet(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// convert to dto output
	output := ResponsePublicGreetDTO{
		RequestID: requestID,
		Title:     result.Title,
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// send result response
	_, err = w.Write(outputJson)
	if err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
