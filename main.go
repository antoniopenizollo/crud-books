package main

import (
	"github.com/antoniopenizollo/crud-books/database"
	"github.com/antoniopenizollo/crud-books/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}