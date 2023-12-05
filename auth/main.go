package main

import (
	"auth/internals/models/server"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// models.CreateSessionManager()

func main() {
	s := server.NewServer()

	port := "8083"

	if err := s.StartServer(port); err != nil {
		fmt.Println(err.Error())
	}
}
