package todoutility

import (
	database "gofiber/database"
	todo "gofiber/models"

	"github.com/gofiber/fiber/v2"
)

//<-- Get ------------------------------------------- -->
//Get all todos
func GetTodos(c *fiber.Ctx) error {
    var todos []todo.Todo
    result := database.DB.Find(&todos)
    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }
    return c.Status(200).JSON(&todos)
}

//Get todo by id
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo []todo.Todo
	database.DB.First(&todo, id)
	if todo == nil {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&todo)
}

//<-- Post ------------------------------------------- -->
//Add new todo
func AddTodo(c *fiber.Ctx) error {
	todo := new(todo.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.SendStatus(400)
	}
	database.DB.Create(&todo)
	return c.Status(200).JSON(&todo)
}


//<-- Put ------------------------------------------- -->
//Update todo
func UpdateTodo(c *fiber.Ctx) error {
	todo := new(todo.Todo)
	id := c.Params("id")
	if err := c.BodyParser(todo); err != nil {
        return c.Status(503).SendString(err.Error())
    }
	database.DB.Where("id = ?", id).Updates(&todo)
	return c.Status(200).JSON(&todo)
}

//<-- Delete ------------------------------------------- -->
//Delete todo
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo []todo.Todo
	database.DB.First(&todo, id)
	if todo == nil {
		return c.SendStatus(404)
	}
	database.DB.Delete(&todo)
	return c.Status(200).JSON(&todo)
}
