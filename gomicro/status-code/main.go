package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Get")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusCreated).Send("Post")
	})

	app.Listen(8080)
}
