package main

import (
	"fmt"
	"log"
)

type application struct {
	appName string
	server  server
	debug   *log.Logger
	infoLog *log.Logger
}
type server struct {
	host string
	port string
	url  string
}

func main() {
	fmt.Println("hello")

}
