package seeder

import "gorm.io/gorm"

type UserSeeder struct {
	DB *gorm.DB
}

func (s *UserSeeder) Seed(count int) {
	// Your seeder here
}
