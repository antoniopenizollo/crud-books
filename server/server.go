package server

import (
	"github.com/antoniopenizollo/crud-books/configs"
	"github.com/antoniopenizollo/crud-books/server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   configs.GetServerPort(),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}