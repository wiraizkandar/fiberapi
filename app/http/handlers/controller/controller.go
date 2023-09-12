package controller

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("Welcome to " + os.Getenv("APP_NAME"))
}
