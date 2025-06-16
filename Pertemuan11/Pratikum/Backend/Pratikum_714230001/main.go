package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"strings"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("gagagl memuat file.env", err)
	}

}

func main() {
	app := fiber.New()

	//logging request di terminal
	app.Use(logger.New())

	//Basic Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	//Router Mahasiswa
	router.SetupRoutes(app)

	//Handler 404
	app.Use(func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"message": "Endpoint tidak ditemukan",
		})
	})

	//Baca PORT yang ada di .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //default port kalau tidak ada di .env
	}

	//untuk log cek konek di port mana
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Errror starting server: %v", err)
	} //Koneksi terputus
}

