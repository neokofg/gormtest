package main

import (
	"gormtest/database"
	"gormtest/routes"
)

func main() {
	db := database.InitializeDatabase()
	routes.InitializeRoutes(db)
}
