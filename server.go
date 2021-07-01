package main

import (
	"log"
	"net/http"
	"os"
)

type Server struct {
}

const port = "8080"

func (server *Server) Start() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started at http://" + hostname + ":" + port)
	err = http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
