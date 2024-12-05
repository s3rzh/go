package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	server := &http.Server{Addr: ":8080"}
	server.RegisterOnShutdown(doCleanup)

	go func() {
		<-ctx.Done()
		log.Println("Got shutdown signal.")
		<-time.After(5 * time.Second)

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Error while stopping HTTP listener: %s", err)
		}
	}()

	log.Fatal(server.ListenAndServe())
}

func doCleanup() {
	log.Println("Cleanup starting...")
	time.Sleep(5 * time.Second)
	log.Println("Cleanup complete.")
}
