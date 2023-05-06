package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// List of app routes
func AppRoutes(r *gin.Engine, db *gorm.DB) {
	apiRoutes(r, db)
	// Write your routes group here ...

}
