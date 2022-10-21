package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "Walk the dog", Completed: false},
	{Id: 2, Name: "Walk the cat", Completed: false},
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello world")
	})

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
