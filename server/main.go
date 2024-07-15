package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	db, err := sqlx.Connect("postgres", "host=localhost port=5433 user=postgres password=sercret dbname=finance_db sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %\nv", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		var result int
		err = db.Get(&result, "SELECT 1")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message":  "Hello, World!",
			"database": "available",
		})
	})
	log.Fatal(app.Listen(":4000"))
}
