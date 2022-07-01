package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context, m string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": m,
	})
}

func InternalServerError(c *gin.Context, m string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusInternalServerError,
		"message": m,
	})
}
