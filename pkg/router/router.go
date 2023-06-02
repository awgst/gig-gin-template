package router

import (
	"gig-gin-template/pkg/database"
	"gig-gin-template/src/routes"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Run router
func Run(db database.Connection) http.Handler {
	path, _ := filepath.Abs("./pkg/templates/index.html")
	r := gin.Default()
	r.LoadHTMLFiles(path)
	routes.AppRoutes(r, db)

	return r
}
