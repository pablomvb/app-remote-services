package main

import (
	"context"
	"log"
	"net/http"
	"path"
	"time"
)

const version = "0.0.1"

type CommandDetail struct {
	name string
	path string
}

type JobDetail struct {
	id          int8
	finished    bool
	command     string
	title       string
	progress    int8
	duration    int8 // time in seconds
	transcurred int8 // time in seconds
}

func initCommands() []CommandDetail {
	commands := []CommandDetail{
		{"command1", path.Join("command1", "command1")},
		{"command2", path.Join("command2", "command2")},
		{"command3", path.Join("command3", "command3")},
	}
	return commands
}

func serviceStopServer(server *http.Server, signalOff chan bool) {
	log.Println("Waiting for signal to stop server")
	state_off := <-signalOff
	for state_off != true {
		log.Println("Waiting for signal to stop server")
		state_off = <-signalOff
	}
	log.Println("Signal received to stop server")
	time.Sleep(time.Millisecond * 500)
	log.Println("Stopping server...")
	server.Shutdown(context.Background())
}

func main() {
	log.Println("API Server v", version)
	initCommands()
	log.Println("Initializing server...")

	var server http.Server
	signalStopServer := make(chan bool)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Hello, World!"}`))
		log.Println(r)
		log.Println("Hello, World!")
	})

	http.HandleFunc("/service/stop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Shutdown Server init")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Shutdown Server"}`))
		log.Println("Shutdown Server send response")
		signalStopServer <- true
		log.Println("Shutdown Server send signal")
	})

	go serviceStopServer(&server, signalStopServer)

	server.Addr = ":8080"
	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
