package main

import (
	"fmt"
	"tdd/server"
)

func main() {
	ready := make(chan bool)
	server := server.StartServer("localhost", "8080", ready)
	<-ready
	fmt.Println("Servidor escuchando en", server.Addr)
	select {}
}
