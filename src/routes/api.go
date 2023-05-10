package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Api routes list
func apiRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})
	// Define your routes here
}
