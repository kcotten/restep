package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	a "restep/app"
	"syscall"
	"time"
)

// Listen and serve the endpoint
func runHttpEndpoint(a *a.App) {
	setupCloseHandler()

	srv := &http.Server{
		Handler: a.Router,
		Addr:    ":8000",
		// Enforce timeouts for server
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Handle exit signal
func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("\nCtrl+C pressed in Terminal, exiting...")
		os.Exit(0)
	}()
}
