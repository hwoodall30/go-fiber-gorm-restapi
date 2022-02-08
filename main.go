package main

import (
	todocontroller "gofiber/controllers"
	database "gofiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()
    database.Connect()
    app.Use(logger.New())
    app.Static("/", "./templates/index.html")
    todocontroller.Routes(app)
    app.Listen(":3000")
}