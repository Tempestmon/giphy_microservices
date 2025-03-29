package handlers

import (
	"fmt"
	"giphy_microservices/internal/domain"
	"giphy_microservices/internal/infrastructure/redis"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Posts       []domain.Post
	RedisClient redis.RedisClient
}

func (h *PostHandler) GetPost(c *gin.Context) {
	key := c.Param("key")
	value, err := h.RedisClient.GetKey(string(key))
	fmt.Printf("error", err)
	if err != nil {
		c.JSON(404, "Key not found")
		return
	}
	c.JSON(200, value)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	value, err := io.ReadAll(c.Request.Body)
	slog.Debug("Reading body")
	if err != nil {
		slog.Warn("Got error reading body")
		c.JSON(405, "Return body")
		return
	}
	keyValue, keyError := h.RedisClient.SetKey("kek", string(value))
	if keyError != nil {
		c.JSON(400, "Internal Error")
		return
	}
	c.JSON(200, keyValue)
}
