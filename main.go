package main

import (
	"task_activity/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API corriendo y DB conectada")
	})

	app.Listen("Servidor escuchando en el puerto :3000")
}
