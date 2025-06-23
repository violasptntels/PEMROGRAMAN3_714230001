package router

import (
	"inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.Middlewares("admin"))
	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", middleware.Middlewares("admin"), handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.GetMahasiswaByNPM)
	api.Post("/mahasiswa", middleware.Middlewares("admin"), handler.CreateMahasiswa)
	api.Put("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.UpdateMahasiswa)
	api.Delete("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.DeleteMahasiswa)
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Get("/docs/*", swagger.HandlerDefault)
}
