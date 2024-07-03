package main

import (
	"gofiber/middlewares"
	"gofiber/posts"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := fiber.Config{
		AppName:           "Learn Golang with Fiber",
		EnablePrintRoutes: true,
		Immutable:         false,
		ServerHeader:      "Server 1",
	}

	app := fiber.New(config)
	app.Use(middlewares.AddRequestId, logger.New())

	api := app.Group("/api").Name("API ")

	apiV1 := api.Group("/v1").Name("v1 ")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}).Name("Hello World")

	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}).Name("Hello World")

	apiV1.Get("/posts", posts.GetAllPosts).Name("Get All Posts")
	apiV1.Post("/posts", posts.CreatePost).Name("Create New Post")
	apiV1.Put("/posts/:id<int>", posts.UpdatePost).Name("Update Post")
	apiV1.Delete("/posts/:id<int>", posts.DeletePostById).Name("Delete Post")
	apiV1.Get("/author", posts.GetPostAuthor).Name("Get Post Author")
	apiV1.Get("/posts/:id<int>", posts.GetPostById).Name("Get Post By Id")
	apiV1.Get("/posts/search", posts.SearchPosts).Name("Search Post")

	log.Fatal(app.Listen(":8080"))

}
