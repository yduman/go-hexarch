package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yduman/go-hexarch/internal/adapters/inbound/http"
	"github.com/yduman/go-hexarch/internal/adapters/outbound/database"
	"github.com/yduman/go-hexarch/internal/application/service"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := database.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	app := fiber.New()
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)

	log.Println("Server started at :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
