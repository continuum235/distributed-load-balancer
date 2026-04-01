package server

import (
	"net/http"

	"github.com/continuum235/distributed-load-balancer/internal/handlers"
	"github.com/continuum235/distributed-load-balancer/internal/middleware"
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
