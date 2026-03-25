package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	e.GET("/canvas/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":  id,
			"url": "/canvas/" + id,
		})
	})

	port := os.Getenv("CANVAS_PORT")
	if port == "" {
		port = "18793"
	}

	log.Printf("Starting canvas service on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start canvas service: %v", err)
	}
}
