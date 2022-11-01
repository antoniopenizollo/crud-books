package main

import (
	"crud-books/database"
	"crud-books/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}