package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/continuum235/distributed-load-balancer/internal/server"
)

func main() {
	fmt.Println("Starting server...")
	helloSrv, adminSrv := server.NewServers()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go helloSrv.ListenAndServe()
	go adminSrv.ListenAndServe()

	sig := <-sigs
	fmt.Println(sig)

	server.ShutdownServers(ctx, helloSrv, adminSrv)
	fmt.Println("services has stopped")
}
