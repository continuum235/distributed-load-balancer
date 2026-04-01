package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func NewServers() (*http.Server, *http.Server) {
	helloMux := SetupHelloRoutes()
	adminMux := SetupAdminRoutes()

	helloSrv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      helloMux,
	}

	adminSrv := &http.Server{
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      adminMux,
	}

	return helloSrv, adminSrv
}

func ShutdownServers(ctx context.Context, servers ...*http.Server) {
	for _, srv := range servers {
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("shutdown error:", err)
		}
	}
}