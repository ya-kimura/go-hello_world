package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Friends! Are you ready?! - Welcome to the Golang Universe!!!")
	})

	app.Listen(8080)
}
