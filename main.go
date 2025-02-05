package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	name := "Sasank Tanikella."
	fmt.Println("Hello!", name)
	app := fiber.New() // Create a new Fiber instance

	todos := []Todo{} // Create a slice of Todo structs

	// create a route on your app, in this case a GET request to "/api/todos", that returns the todos slice
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos) // Return the todos slice
	})

	// create a route on your app, in this case a POST request to "/api/todos", that adds a new Todo to the todos slice
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // Create a new Todo struct {id:0, completed:false, body:""} (default values) the body is coming from the request body
		if err:= c.BodyParser(&todo);  err != nil { // Parse the body of the request into the todo struct
			return err
		}
		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"}) // Return an error response if the body is empty
		}
		todo.ID = len(todos) + 1 // Set the ID of the todo to the length of the todos slice + 1
		todos = append(todos, *todo) // Append the todo to the todos slice

		return c.Status(201).JSON(todo) // Return the created todo
	})

	// create a route on your app, in this case a PATCH request to "/api/todos/:id", that updates a Todo in the todos slice
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id") // Get the id parameter from the request

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true // Set the completed field of the todo to true
				return c.Status(200).JSON(todos[i]) // Return the updated todo
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"}) // Return an error response if the todo is not found
	})

	// create a route on your app, in this case a DELETE request to "/api/todos/:id", that deletes a Todo from the todos slice
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id") // Get the id parameter from the request

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i], todos[i+1:]...) // Remove the todo from the todos slice
				return c.Status(200).JSON(fiber.Map{"msg": "Todo deleted"}) // Return a success response
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"}) // Return an error response if the todo is not found
	})

	// Create a route on your app, in this case a GET request to "/test", that returns a JSON response
	app.Get("/test", func(c *fiber.Ctx) error { // c is the context that is passed to the handler, it contains information about the request and the response
		return c.Status(200).JSON(fiber.Map{"msg": "test complete"})
	})

	log.Fatal(app.Listen(":8080")) // Listen on port 8080
}
