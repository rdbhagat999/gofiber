package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var postList []Post = []Post{
	{Id: 1, Title: "First post"},
	{Id: 2, Title: "Second post"},
	{Id: 3, Title: "Third post"},
}

func getAllPosts(c *fiber.Ctx) error {
	return c.JSON(postList)
}

func createPost(c *fiber.Ctx) error {
	var post Post

	err := c.BodyParser(&post)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}

	post.Id = len(postList) + 1

	postList = append(postList, post)

	return c.JSON(post)
}

func getPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
	}

	for _, item := range postList {
		if item.Id == id {
			return c.JSON(item)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Post not found")
}

func main() {
	config := fiber.Config{
		AppName: "Learn Golang with Fiber",
	}

	app := fiber.New(config)

	api := app.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	apiV1.Get("/posts", getAllPosts)
	apiV1.Post("/posts", createPost)
	apiV1.Get("/posts/:id", getPostById)

	log.Fatal(app.Listen(":8080"))

}
