package server

import (
	"dist_load_balancer/internal/handlers"
	"dist_load_balancer/internal/middleware"
	"net/http"
)

func SetupHelloRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.Hello)

	return middleware.Middleware{Mux: mux}
}

func SetupAdminRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)

	return middleware.Middleware{Mux: mux}
}