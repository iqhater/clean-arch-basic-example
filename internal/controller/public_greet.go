package controller

import (
	srv "clean_arch_super_simple_example/internal/usecase"
	"clean_arch_super_simple_example/pkg"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const contextNameKey contextKey = "name"

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
	reqID, ok := req.Context().Value(pkg.RequestIDKey).(uuid.UUID)
	if !ok || reqID == uuid.Nil {
		fmt.Fprintln(w, "this a request without requestID")
		return
	}
	fmt.Fprintf(w, "request id is %s\n", reqID)

	// name and method request validations are located in separate validateRequest middleware
	// get id and name params from context
	name := req.Context().Value(contextNameKey).(string)

	// w.Header().Add("Content-Type", "application/json")

	output := g.greetService.GetGreet(name)

	// send result response
	w.Write([]byte(output))
}
