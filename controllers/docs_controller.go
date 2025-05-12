package controllers


import (
	"github.com/gofiber/fiber/v3"
)


func GetDocs(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
