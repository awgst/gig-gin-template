package main

import (
	"fmt"
	"gig-gin-template/database/seeder"
)

func main() {
	seeder.Execute()
	fmt.Println("Database seeded!")
}
