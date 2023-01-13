package handlers

import (
	" hery-ciaputra/demo-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	productService ProductService
}

func New(ps ProductService) *Handler {
	return &Handler{
		productService: ps,
	}
}

func (h *Handler) handler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	result := db.Get().First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR", // to give specific which 404 error
			"statusCode": http.StatusNotFound,
			"message":    "product not found",
		})
		return
	}
	c.Set("user", "Hery")
	c.JSON(http.StatusOK, product)
}
