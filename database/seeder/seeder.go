package seeder

import (
	"gorm.io/gorm"
)

type SeederInterface interface {
	Seed(db *gorm.DB)
}

func Execute() {
	// Uncomment this line
	// db := database.Connect()

	// Your seeder here
}
