package main

var server Server
var routes Routes

func main() {
	routes.Register()
	server.Start()
}
