package main

var (
	server Server
	routes Routes
)

func main() {
	routes.Register()
	server.Start()
}
