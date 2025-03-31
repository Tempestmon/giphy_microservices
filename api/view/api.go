package view

import (
	"api/config"
	"api/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.ServerConfig
	server *gin.Engine
}

func (s *Server) Start() error {
	return s.server.Run(fmt.Sprintf("%v:%v", s.config.IP, s.config.Port))
}

func NewServer(config config.ServerConfig) *Server {
	router := gin.Default()
	{
		v1 := router.Group("/v1")
		v1.GET("/meme/:id", handlers.GetMeme)
		v1.POST("/meme", handlers.SetMeme)
	}
	return &Server{
		config: config,
		server: router,
	}
}
