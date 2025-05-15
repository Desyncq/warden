package main

import (
	"github.com/Desyncq/warden/controllers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	registerRoutes(app)

	app.Listen(":3000")
}

func registerRoutes(app *fiber.App) {
	// Serve Frontend
	app.Get("/*", static.New("./web"))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("./web/index.html")
	})

	docs := app.Group("/docs")
	docs.Get("/", controllers.GetDocs)

	// Admin Routes
	admin := app.Group("/admin")
	adminController := controllers.NewAdminController()
	admin.Get("/signup", adminController.AdminSignup)
	admin.Get("/login", adminController.AdminLogin)
	admin.Get("/logout", adminController.AdminLogout)

}
