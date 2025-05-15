package controllers

import (
	"github.com/gofiber/fiber/v3"
)


type AdminController struct {}


func (a *AdminController) AdminSignup(c fiber.Ctx) error {
	return c.SendString("Admin Signup")
}

func (a *AdminController) AdminLogin(c fiber.Ctx) error {
	return nil
}

func (a *AdminController) AdminLogout(c fiber.Ctx) error {
	return nil
}


func NewAdminController() *AdminController {
	return &AdminController{}
}
