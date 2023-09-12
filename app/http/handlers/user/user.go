package user

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	return c.SendString("This is Get Users")
}

func GetUserById(c *fiber.Ctx) error {
	return c.SendString("This user id is " + c.Params("id"))
}
