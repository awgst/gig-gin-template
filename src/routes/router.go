package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// List of app routes
func AppRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	// Define your routes group here ...
	apiRoutes(r, db)

}
