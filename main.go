package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CustomError(errCode int, message string) error {
	customErr := fiber.NewError(errCode, message)
	return customErr
}

func main() {
	config := fiber.Config{
		AppName:           "Learn Golang with Fiber",
		EnablePrintRoutes: true,
		Immutable:         false,
		ServerHeader:      "Server 1",
	}

	app := fiber.New(config)

	api := app.Group("/api").Name("API ")

	apiV1 := api.Group("/v1").Name("v1 ")

	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}).Name("Hello World")

	apiV1.Get("/posts", getAllPosts).Name("Get All Posts")
	apiV1.Post("/posts", createPost).Name("Create New Post")
	apiV1.Get("/author", getPostAuthor).Name("Get Post Author")
	apiV1.Get("/posts/:id", getPostById).Name("Get Post By Id")

	log.Fatal(app.Listen(":8080"))

}
