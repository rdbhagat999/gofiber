package middlewares

import (
	"log/slog"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/gofiber/fiber/v2"
)

func AddRequestId(c *fiber.Ctx) error {
	uid, _ := guid.NewV4()
	c.Request().Header.Add("Request-ID", uid.String())
	return c.Next()
}

func RequestLogger(c *fiber.Ctx) error {
	reqId := c.Request().Header.Peek("Request-ID")
	reqIdString := string(reqId)
	slog.Info("request", "method", c.Method(), "path", c.Path(), "params", c.AllParams(), "reqId", reqIdString, "query", c.Query("q"))
	return c.Next()
}
