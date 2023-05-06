package main

import (
	"gig-gin-template/pkg/database"
	"gig-gin-template/pkg/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router.Run(database.Connect())
}
