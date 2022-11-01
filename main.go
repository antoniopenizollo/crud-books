package main

import (
	"github.com/antoniopenizollo/crud-books/configs"
	"github.com/antoniopenizollo/crud-books/database"
	"github.com/antoniopenizollo/crud-books/server"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(any(err))
	}

	database.StartDB()

	server := server.NewServer()

	server.Run()
}