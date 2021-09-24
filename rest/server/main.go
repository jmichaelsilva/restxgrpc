package main

import (
	"github.com/gofiber/fiber/v2"
)

type Image struct {
	Data string `json:"data"`
}

func main() {
	app := fiber.New()

	app.Post("/image", func(c *fiber.Ctx) error {
		var img Image
		if err := c.BodyParser(&img); err != nil {
			return c.JSON("ERROR")
		}

		return c.JSON("OK")
	})

	app.Listen(":3000")
}
