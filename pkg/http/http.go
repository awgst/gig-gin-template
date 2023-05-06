package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Helper for return a json response
func Json(c *gin.Context, message string, data interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	c.JSON(code, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}
