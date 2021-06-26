package main

import (
	"log"
	"net/http"
)

type Server struct {
}

func (server *Server) Start() {
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
