package internal

import (
	ctrl "clean_arch_super_simple_example/internal/controller"
	infra "clean_arch_super_simple_example/internal/infrastructure"
	srv "clean_arch_super_simple_example/internal/usecase"
	"clean_arch_super_simple_example/pkg"
	"clean_arch_super_simple_example/pkg/logger"
	"fmt"
	"net/http"
)

// config
const DB_FILENAME = "db.txt"
const HTTP_PORT = "8080"

// init
func Run() {

	// repo/db init
	repo := infra.NewGreetRepository(DB_FILENAME)
	repoWithCounter := infra.NewPrivateGreetRepository(DB_FILENAME)

	// service init
	service := srv.NewGreetService(repo)
	serviceWithCounter := srv.NewPrivateGreetService(repoWithCounter)

	// controller init
	controller := ctrl.NewGreetController(service)
	controllerWithCounter := ctrl.NewPrivateGreetController(serviceWithCounter)

	// router init
	http.HandleFunc("/greet", logger.Log(ctrl.ValidateRequest(pkg.RequestID(controller.GreetHandler))))
	http.HandleFunc("/greet/{id}", logger.Log(ctrl.ValidateRequest(controllerWithCounter.GreetPrivateHandler)))

	// server init
	fmt.Printf("🌐 Clean Arch Example API server started on port: %s\n", HTTP_PORT)
	http.ListenAndServe(":"+HTTP_PORT, nil)
}
