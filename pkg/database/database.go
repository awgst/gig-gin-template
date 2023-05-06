package database

import (
	"fmt"
	"gig-gin-template/pkg/env"
	"gig-gin-template/src/config"
	"log"

	"golang.org/x/exp/slices"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// All valid drivers
var validDrivers = []string{
	"mysql",
}

// Database driver that get from .env
var driver = env.Get("DB_DRIVER", "")

// Dialector for open connection based on driver
var dialectors = map[string]gorm.Dialector{
	"mysql": mysql.Open(getDbDsn()),
}

// Connect and return *gorm.DB connection
func Connect() *gorm.DB {
	dialector := dialectors[driver]
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error", err)
	}

	fmt.Println("Database connected successfully")
	return db
}

// Convert map database configuration into single line database dsn
func getDbDsn() string {
	if slices.Contains(validDrivers, driver) {
		log.Fatalf("Driver %s is unavailable", driver)
	}

	config := config.Database()[driver]

	dbDsn := map[string]string{
		"mysql": fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config["username"],
			config["password"],
			config["host"],
			config["port"],
			config["database"],
		),
	}

	return dbDsn[driver]
}
