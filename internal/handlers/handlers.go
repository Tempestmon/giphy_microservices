package handlers

import (
	"fmt"
	"giphy_microservices/internal/domain"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	posts []domain.Post
}

func (h *PostHandler) GetPost(c *gin.Context) {
	fmt.Println("Post")
}
