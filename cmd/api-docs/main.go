package main

import (
	"clean_arch_basic_example/pkg/logger"
	"fmt"
	"log"
	"net/http"
	"time"
)

// OpenAPI docs
func main() {

	// serve api folder for openapi docs
	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("./api"))))

	// redirect to from / to /api/redoc.html
	mux.HandleFunc("/", logger.Log(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/redoc.html", http.StatusMovedPermanently)
	}))

	const HTTP_PORT = ":3200"

	fmt.Printf("📄 Clean Arch Example API Docs server started on port: %s\n", HTTP_PORT)

	// best practice to use timeout
	server := &http.Server{
		Addr:              HTTP_PORT,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}
	log.Fatal(server.ListenAndServe())
}
