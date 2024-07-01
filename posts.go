package main

import (
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
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
		customErr := CustomError(fiber.StatusBadRequest, "Invalid data")
		return customErr
	}

	post.Id = len(postList) + 1

	postList = append(postList, post)

	return c.JSON(post)
}

func getPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		customErr := CustomError(fiber.StatusBadRequest, "Invalid id")
		return customErr
	}

	for _, item := range postList {
		if item.Id == id {
			return c.JSON(item)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Post not found")
}

func getPostAuthor(c *fiber.Ctx) error {
	author := struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{
		Id:   1,
		Name: "John",
	}

	// result, _ := json.Marshal(author)
	// c.Response().Header.Add("Content-Type", "application/json")
	// return c.Send(result)

	return c.JSON(author)
}
