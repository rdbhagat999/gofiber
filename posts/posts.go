package posts

import (
	"gofiber/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EditPost struct {
	Title string `json:"title"`
}

var postList []Post = []Post{
	{Id: 1, Title: "First post"},
	{Id: 2, Title: "Second post"},
	{Id: 3, Title: "Third post"},
}

func GetAllPosts(c *fiber.Ctx) error {
	return c.JSON(postList)
}

func CreatePost(c *fiber.Ctx) error {
	var post Post

	err := c.BodyParser(&post)

	if err != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid data")
		return customErr
	}

	post.Id = len(postList) + 1

	postList = append(postList, post)

	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {

	var post EditPost

	id, idErr := c.ParamsInt("id")

	if idErr != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid id")
		return customErr
	}

	err := c.BodyParser(&post)

	if err != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid request")
		return customErr
	}

	for i, item := range postList {

		if item.Id == id {
			found := &postList[i]
			found.Title = post.Title
		}

	}

	return c.JSON(post)
}

func GetPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid id")
		return customErr
	}

	for _, item := range postList {
		if item.Id == id {
			return c.JSON(item)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Post not found")
}

func DeletePostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		// return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid id")
		return customErr
	}

	for index, item := range postList {
		if item.Id == id {
			postList = append(postList[:index], postList[index+1:]...)

		}
	}

	return c.Status(fiber.StatusOK).SendString("Post deleted")
}

func SearchPosts(c *fiber.Ctx) error {
	query := c.Query("q")

	if query == "" {
		customErr := utils.CustomError(fiber.StatusBadRequest, "Invalid search request")
		return customErr
	}

	var searchResults = []Post{}

	for _, item := range postList {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(query)) {
			searchResults = append(searchResults, item)
		}
	}

	return c.JSON(searchResults)
}

func GetPostAuthor(c *fiber.Ctx) error {
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
