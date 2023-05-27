package routes

import (
	"gig-gin-template/pkg/database"

	"github.com/gin-gonic/gin"
)

// List of app routes
func AppRoutes(r *gin.Engine, db database.Connection) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	// Define your routes group here ...
	apiRoutes(r, db)
}
