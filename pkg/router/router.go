package router

import (
	"fmt"
	"gig-gin-template/pkg/env"
	"gig-gin-template/src/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Run router
func Run(db *gorm.DB) {
	ginMode := env.Get("GIN_MODE", gin.DebugMode)
	if ginMode == gin.ReleaseMode {
		gin.SetMode(ginMode)
	}
	r := gin.Default()
	routes.AppRoutes(r, db)

	port := fmt.Sprintf(":%s", env.Get("APP_PORT", "8080"))

	r.Run(port)
}
