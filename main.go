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
	SetupApiV1(app)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}

func SetupApiV1(app *fiber.App) {
	v1 := app.Group("/v1")

	SetupTodosRoutes(v1)
}

func SetupTodosRoutes(grp fiber.Router) {
	todosRoutes := grp.Group("/todos")
	todosRoutes.Get("/", GetTodos)
	todosRoutes.Post("/", CreateTodo)
}

func GetTodos(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}

func CreateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	todo := &Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	ctx.Status(fiber.StatusCreated).JSON(todo)
}
