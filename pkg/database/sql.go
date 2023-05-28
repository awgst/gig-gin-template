package database

import (
	"database/sql"
	"fmt"
	conf "gig-gin-template/src/config"
	"log"
	// Uncomment the code based on your database driver
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

func ConnectSql(config ...conf.DatabaseConfig) *sql.DB {
	dbDsn := getDbDsn()
	if len(config) > 0 {
		dbDsn = config[0].DbDsn()
	}

	// Open database connection
	db, err := sql.Open(driver, dbDsn)
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	fmt.Println("Database connected successfully")
	return db
}
