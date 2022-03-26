package server

import (
	"auth-api/server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s Server) Run() {
	route := routes.ConfigRoutes(s.server)

	log.Println("Server is running on port: ", s.port)
	log.Fatalln(route.Run(":" + s.port))
}
