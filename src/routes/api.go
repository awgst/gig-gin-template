package routes

import (
	"gig-gin-template/pkg/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Api routes list
func apiRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		http.Json(c, "GIG help faster Go project development. More informations --> https://github.com/awgst/gig", nil)
	})
}
