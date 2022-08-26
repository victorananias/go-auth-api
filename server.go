package main

import (
	"log"
	"net/http"
)

type Server struct {
}

const PORT = "8080"

func (server *Server) Start() {
	log.Println("Server started at :" + PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	log.Fatal(err)
}
