package routes

import (
	"gig-gin-template/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List of app routes
func AppRoutes(r *gin.Engine, db database.Connection) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	// Define your routes here
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})
}
