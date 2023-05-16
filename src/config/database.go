package config

import (
	"gig-gin-template/pkg/env"
	"gig-gin-template/src/config/database"
)

type DatabaseConfig interface {
	DbDsn() string
}

type conf map[string]DatabaseConfig

// Return type conf of database configurations from each driver
func Database() conf {
	return conf{
		"mysql": &database.MySql{
			Host:     env.Get("DB_HOST", "localhost"),
			Port:     env.Get("DB_PORT", "3306"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			Database: env.Get("DB_DATABASE", "gig"),
		},
		"postgresql": &database.PostgreSql{
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
