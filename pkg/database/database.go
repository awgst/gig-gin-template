package database

import (
	"database/sql"
	"gig-gin-template/pkg/common/slices"
	"gig-gin-template/pkg/env"
	"gig-gin-template/src/config"
	"log"

	"gorm.io/gorm"
)

// Database interface
type Connection struct {
	SQL  *sql.DB
	Gorm *gorm.DB
}

// All valid drivers
var validDrivers = []string{
	"mysql",
	"postgres",
}

// Database driver that get from .env
var driver = env.Get("DB_DRIVER", "")

// Convert map database configuration into single line database dsn
func getDbDsn() string {
	if !slices.Contains(validDrivers, driver) {
		log.Fatalf("Driver %s is unavailable", driver)
	}

	config := config.Database()[driver]

	return config.DbDsn()
}
