package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	srv "github.com/iqhater/clean-arch-basic-example/internal/usecase"
	mid "github.com/iqhater/pkg/middleware"
)

type contextKey string

const contextNameKey contextKey = "name"

type ResponsePublicGreetDTO struct {
	RequestID string `json:"request_id"`
	Title     string `json:"greeting"`
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
	requestID := mid.IDFromContext(req.Context())

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
