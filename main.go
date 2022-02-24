package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 👌")
	})

	app.Get("/narnia/*", func(c *fiber.Ctx) error {
		message := fmt.Sprintf("👏🏼 %s", c.Params("*"))
		return c.SendString(message)
	})

	app.Listen(":3002")
}
