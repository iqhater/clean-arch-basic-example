package internal

import (
	ctrl "clean_arch_basic_example/internal/controller"
	infra "clean_arch_basic_example/internal/infrastructure"
	srv "clean_arch_basic_example/internal/usecase"
	"clean_arch_basic_example/pkg"
	"clean_arch_basic_example/pkg/logger"
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/greet", logger.Log(ctrl.ValidateRequest(pkg.RequestID(controller.GreetHandler))))
	http.HandleFunc("/greet/{id}", logger.Log(ctrl.ValidateRequest(controllerWithCounter.GreetPrivateHandler)))

	// server init
	fmt.Printf("🌐 Clean Arch Example API server started on port: %s\n", cfg.HTTP_PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTP_PORT, nil))
}
