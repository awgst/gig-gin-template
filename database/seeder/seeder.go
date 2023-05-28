package seeder

import (
	"context"
	"gig-gin-template/pkg/database"
	"gig-gin-template/pkg/env"
	conf "gig-gin-template/src/config/database"
)

type SeederInterface interface {
	Seed(ctx context.Context, count int)
}

func Execute() {
	ctx := context.Background()
	conn := database.Connection{
		SQL: database.ConnectSql(&conf.MySql{
			Host:     env.Get("FORWARD_DB_HOST", "localhost"),
			Port:     env.Get("FORWARD_DB_PORT", "3306"),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			Database: env.Get("DB_DATABASE", "gig"),
		}),
	}
	db := conn.SQL
	// Your seeder here
	userSeeder := UserSeeder{DB: db}
	userSeeder.Seed(ctx, 1)
}
