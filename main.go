package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/Desyncq/warden/controllers"
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
	app.Get("/", func (c fiber.Ctx) error {
		return c.SendFile("./web/index.html")
	})

	docs := app.Group("/docs")
	docs.Get("/", controllers.GetDocs)

	// Admin Routes
	admin := app.Group("/admin")
	admin.Get("/signup", controllers.AdminSignup)
	admin.Get("/login", controllers.AdminLogin)
	admin.Get("/logout", controllers.AdminLogout)


}
