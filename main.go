package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/phillip-d-shields/go-fiber-crm/database"
	"github.com/phillip-d-shields/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead/", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection successfully opened.")
	err = database.DBConn.AutoMigrate(&lead.Lead{})
	if err != nil {
		panic("Failed to migrate database!")
	}
	fmt.Println("Database migrated.")
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")
}
