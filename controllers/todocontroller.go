package todocontroller

import (
	todoutility "gofiber/utilities"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/api/todos", todoutility.GetTodos)
	app.Get("/api/todo/:id", todoutility.GetTodo)
	app.Post("/api/addtodo", todoutility.AddTodo)
	app.Put("/api/updatetodo/:id", todoutility.UpdateTodo)
	app.Delete("/api/deletetodo/:id", todoutility.DeleteTodo)
}