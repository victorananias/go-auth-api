package main

import "net/http"

type Routes struct {
}

func (routes *Routes) Register() {
	http.HandleFunc("/", List)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
}
