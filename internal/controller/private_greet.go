package controller

import (
	srv "clean_arch_basic_example/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const contextIDKey contextKey = "id"

type ResponsePrivateGreetDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"greeting"`
	Counter int    `json:"total_greetings"`
}

// controller/handler
type GreetPrivateController struct {
	greetPrivateService *srv.GreetPrivateService
}

// init
func NewPrivateGreetController(greetPrivateService *srv.GreetPrivateService) *GreetPrivateController {
	return &GreetPrivateController{
		greetPrivateService: greetPrivateService,
	}
}

// implementation
func (g *GreetPrivateController) GreetPrivateHandler(w http.ResponseWriter, req *http.Request) {

	// validate and convert id value to string
	id, err := strconv.Atoi(req.PathValue(string(contextIDKey)))
	if err != nil {
		log.Println(err)
		http.Error(w, "'id' value must be integer!", http.StatusBadRequest)
		return
	}

	// name and method request validations are located in separate validateRequest middleware
	// get id and name params from context
	name := req.Context().Value(contextNameKey).(string)

	w.Header().Add("Content-Type", "application/json")

	result, err := g.greetPrivateService.GetGreetingsWithCounter(id, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// convert to dto output
	output := ResponsePrivateGreetDTO{
		ID:      result.ID,
		Title:   result.Title,
		Counter: int(result.Counter),
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// send result response
	w.Write(outputJson)
}
