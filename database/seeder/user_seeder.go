package seeder

import "database/sql"

type UserSeeder struct {
	DB *sql.DB
}

func (s *UserSeeder) Seed(count int) {
	// Your seeder here
}
