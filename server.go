package main

import (
	"log"
	"net/http"
)

type Server struct {
}

const port = "8080"

func (server *Server) Start() {
	log.Println("Server started")
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
