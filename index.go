package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}
type defaultResponse struct {
	Message string
}

func (s *server) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	response := defaultResponse{}

	switch request.Method {
	case "GET":
		responseWriter.WriteHeader(http.StatusOK)
		response.Message = "get called"
		r, _ := json.Marshal(response)
		responseWriter.Write(r)
	case "POST":
		responseWriter.WriteHeader(http.StatusCreated)
		response.Message = "post called"
		r, _ := json.Marshal(response)
		responseWriter.Write(r)
	case "PUT":
		responseWriter.WriteHeader(http.StatusAccepted)
		response.Message = "put called"
		r, _ := json.Marshal(response)
		responseWriter.Write(r)
	case "DELETE":
		responseWriter.WriteHeader(http.StatusOK)
		response.Message = "delete called"
		r, _ := json.Marshal(response)
		responseWriter.Write(r)
	default:
		responseWriter.WriteHeader(http.StatusNotFound)
		response.Message = "not found"
		r, _ := json.Marshal(response)
		responseWriter.Write(r)
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
