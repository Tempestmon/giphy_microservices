package handlers

import (
	"giphy_microservices/internal/domain"
	"giphy_microservices/internal/infrastructure/redis"
	"io"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Posts       []domain.Post
	RedisClient redis.RedisClient
}

func (h *PostHandler) GetPost(c *gin.Context) {
	key := c.Param("key")
	value := h.RedisClient.GetKey(string(key))
	if value.Value != "" {
		c.JSON(200, "Hello")
	} else {
		c.JSON(400, "No keys")
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	value, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(405, "Return body")
	}
	h.RedisClient.SetKey("kek", string(value))
	c.JSON(200, "Ok")
}
