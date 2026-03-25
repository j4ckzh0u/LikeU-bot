package main

import (
	"log"
	"os"

	"ai-coding-assistant/internal/api/middleware"
	"ai-coding-assistant/internal/api/router"
	"ai-coding-assistant/internal/repository"
	"ai-coding-assistant/internal/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := repository.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient := repository.NewRedis(os.Getenv("REDIS_URL"))
	defer redisClient.Close()

	repo := repository.NewRepository(db, redisClient)
	svc := service.NewService(repo)
	h := &router.Handler{
		Service: svc,
	}

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())
	e.Use(middleware.JWT())

	router.SetupRoutes(e, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
