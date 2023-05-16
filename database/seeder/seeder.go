package seeder

import "gig-gin-template/pkg/database"

type SeederInterface interface {
	Seed(count int)
}

func Execute() {
	db := database.Connect()

	// Your seeder here
	userSeeder := UserSeeder{DB: db}
	userSeeder.Seed(1)
}
