package apiservice

import (
	"giphy_microservices/internal/domain"
	"giphy_microservices/internal/handlers"
	"giphy_microservices/internal/infrastructure/redis"

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
	handler := handlers.PostHandler{
		Posts:       []domain.Post{},
		RedisClient: *redis.NewRedisClient("127.0.0.1:6379"),
	}
	{
		v1 := router.Group("/v1")
		v1.GET("/posts/:id", handler.GetPost)
		v1.POST("/post", handler.CreatePost)
	}
	return &Server{
		router:  router,
		address: address,
	}
}
