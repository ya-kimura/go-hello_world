package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Post("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Put("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Delete("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Patch("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Listen(8080)
}
