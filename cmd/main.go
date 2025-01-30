package main

import (
	"net/http"
	"time"

	Server "real-time-forum/internal" // Ensure this import path is correct
)

func main() {
	s := Server.Server{
		HTTP: &http.Server{
			Addr:         ":8080",
			Handler:      Server.Routes(), 
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}

	s.RunServer()
}
