package database

import (
	"database/sql"

	"gorm.io/gorm"
)

// Database interface
type Connection struct {
	SQL  *sql.DB
	Gorm *gorm.DB
}

// All valid drivers
var ValidDrivers = []string{
	"mysql",
	"postgres",
}
