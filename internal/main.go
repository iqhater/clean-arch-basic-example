package internal

import (
	"fmt"
	"log"
	"net/http"

	ctrl "github.com/iqhater/clean-arch-basic-example/internal/controller"
	infra "github.com/iqhater/clean-arch-basic-example/internal/infrastructure"
	srv "github.com/iqhater/clean-arch-basic-example/internal/usecase"
	mid "github.com/iqhater/pkg/middleware"
)

// init
func Run() {

	// init config
	cfg := NewConfig()

	// repo/db init
	repo := infra.NewGreetRepository(cfg.DB_FILENAME)
	repoWithCounter := infra.NewPrivateGreetRepository(cfg.DB_FILENAME)

	// service init
	service := srv.NewGreetService(repo)
	serviceWithCounter := srv.NewPrivateGreetService(repoWithCounter)

	// controller init
	controller := ctrl.NewGreetController(service)
	controllerWithCounter := ctrl.NewPrivateGreetController(serviceWithCounter)

	// router init
	http.HandleFunc("/greet", mid.Log(ctrl.ValidateRequest(mid.RequestID(http.HandlerFunc(controller.GreetHandler)))).ServeHTTP)
	http.HandleFunc("/greet/{id}", mid.Log(ctrl.ValidateRequest(http.HandlerFunc(controllerWithCounter.GreetPrivateHandler))).ServeHTTP)

	// server init
	fmt.Printf("🌐 Clean Arch Example API server started on port: %s\n", cfg.HTTP_PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTP_PORT, nil))
}
