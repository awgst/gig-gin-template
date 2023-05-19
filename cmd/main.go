package main

import (
	"gig-gin-template/pkg/database"
	"gig-gin-template/pkg/router"
)

func main() {
	router.Run(database.Connect())
}
