package main

import (
	"net/http"
)

type Routes struct {
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (routes *Routes) Register() {
	http.HandleFunc("/", setupCors(List))
	http.HandleFunc("/register", setupCors(Register))
	http.HandleFunc("/login", setupCors(Login))
}

func setupCors(handler func(http.ResponseWriter, *http.Request)) HandlerFunc {
	return func(responsewriter http.ResponseWriter, request *http.Request) {
		responsewriter.Header().Set("Access-Control-Allow-Origin", "*")
		responsewriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		responsewriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		handler(responsewriter, request)
	}
}
