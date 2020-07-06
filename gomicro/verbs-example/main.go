package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Get")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Send("Post")
	})

	app.Put("/", func(ctx *fiber.Ctx) {
		ctx.Send("Put")
	})

	app.Delete("/", func(ctx *fiber.Ctx) {
		ctx.Send("Delete")
	})

	app.Patch("/", func(ctx *fiber.Ctx) {
		ctx.Send("Patch")
	})

	app.Listen(8080)
}
