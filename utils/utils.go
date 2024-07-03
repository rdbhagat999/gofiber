package utils

import (
	"github.com/gofiber/fiber/v2"
)

func CustomError(errCode int, message string) error {
	customErr := fiber.NewError(errCode, message)
	return customErr
}
