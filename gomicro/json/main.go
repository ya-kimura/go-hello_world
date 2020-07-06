package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/get-name/:name", func(ctx *fiber.Ctx) {
		name := ctx.Params("name")

		result := map[string]string{
			"name": name,
		}

		ctx.JSON(result)
	})

	app.Listen(8080)
}
