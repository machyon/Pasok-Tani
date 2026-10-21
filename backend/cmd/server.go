package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/machyon/Pasok-Tani/internal/config"
	"github.com/machyon/Pasok-Tani/internal/database"
)

func main() {
    app := fiber.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

    app.Get("/", func(c fiber.Ctx) error {
		connString := cfg.DatabaseURL
		if connString == "" {
			fmt.Fprintf(os.Stderr, "config DATABASE_URL not set\n")
			os.Exit(1)
		}
		ctx := context.Background()
		database.GetConnection(ctx, cfg.DatabaseURL)
      	return c.SendString("Hello World!")
    })

    log.Fatal(app.Listen(":" + cfg.Port))
}
