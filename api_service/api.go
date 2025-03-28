package apiservice

import (
	"giphy_microservices/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	address string
}

func (s *Server) Start() error {
	return s.router.Run(s.address)
}

func NewServer(address string) *Server {
	router := gin.Default()
	handler := handlers.PostHandler{}
	router.GET("/posts/:id", handler.GetPost)

	return &Server{
		router:  router,
		address: address,
	}
}
