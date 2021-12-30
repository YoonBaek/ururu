package main

import (
	"github.com/YoonBaek/ururu-server/board"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	board.Routes(app)
	app.Listen(":3000")
}
