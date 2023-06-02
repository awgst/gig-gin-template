package main

import (
	"fmt"
	"gig-gin-template/pkg/database"
	"gig-gin-template/pkg/env"
	"gig-gin-template/pkg/router"
	config "gig-gin-template/src/config"
	dbconfig "gig-gin-template/src/config/database"
	"log"
	"net/http"
)

func main() {
	// Load .env file
	if err := env.Load(); err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	// Define database connection
	dbConnection := database.Connection{
		// SQL: database.ConnectSql(env.Get("DB_DRIVER"), Database()[env.Get("DB_DRIVER")]),
		// Gorm: database.ConnectGorm(env.Get("DB_DRIVER"), Database()[env.Get("DB_DRIVER")]),
	}

	// Define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.Get("APP_PORT", "8080")),
		Handler: router.Run(dbConnection),
	}

	// Run http server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error running http server ", err)
	}
}

type conf map[string]config.DatabaseConfig

// Return type conf of database configurations from each driver
func Database() conf {
	return conf{
		"mysql": &dbconfig.MySql{
			Host:     env.Get("DB_HOST", "localhost"),
			Port:     env.Get("DB_PORT", "3306"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			Database: env.Get("DB_DATABASE", "gig"),
		},
		"postgres": &dbconfig.PostgreSql{
			Host:     env.Get("DB_HOST", "localhost"),
			Port:     env.Get("DB_PORT", "5432"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			Database: env.Get("DB_DATABASE", "gig"),
			TimeZone: env.Get("DB_TIMEZONE", "Asia/Jakarta"),
		},
		// Your database configuration here
	}
}
