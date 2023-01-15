package main

import (
	"fmt"
	"goMongoDbAPI/app"
	"goMongoDbAPI/configs"
	"goMongoDbAPI/repository"
	"goMongoDbAPI/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	appRoute := fiber.New()
	configs.ConnectDB()
	fmt.Println("Connected to MongoDB!")

	dbClient := configs.GetCollection(configs.DB, "todos")
	fmt.Println("Collection instance created!")

	TodoRepositoryDb := repository.NewTodoRepositoryDB(dbClient)
	fmt.Println("TodoRepositoryDb instance created!")

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}
	fmt.Println("TodoHandler instance created!")

	appRoute.Post("/api/todo", td.CreateTodo)
	fmt.Println("CreateTodo route created!")

	appRoute.Listen(":3000")
	fmt.Println("Server started at port 3000!")
}
