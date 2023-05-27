package routes

import (
	"gig-gin-template/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Api routes list
func apiRoutes(r *gin.Engine, db database.Connection) {
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})
	// Define your routes here
}
