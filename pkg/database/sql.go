package database

import (
	"database/sql"
	"fmt"
	"log"
	// Uncomment the code based on your database driver
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

func ConnectSql() *sql.DB {
	db, err := sql.Open(driver, getDbDsn())
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	fmt.Println("Database connected successfully")
	return db
}
