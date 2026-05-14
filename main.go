package main

import (
	"apiservice/internal/database"
	"apiservice/internal/handler"
	"apiservice/internal/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()

	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	e := echo.New()

	// Auth routes
	authHandler := handler.NewHandler()
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// JDB Handler
	jdbHandler := handler.NewJDBAuthHandler(
		os.Getenv("JDB_AUTH_URL"),
		os.Getenv("JDB_USER_ID"),
		os.Getenv("JDB_SECRET_ID"),
		os.Getenv("JDB_LENDING_URL"),
		os.Getenv("JDB_SECRET_KEY"),
	)

	protected := e.Group("")
	protected.Use(middleware.JWTAuth)
	{
		protected.GET("/me", authHandler.Me)
		protected.POST("/request-lending", jdbHandler.HandleDigitalLending)
	}

	// Health check (public)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "API Service is running",
			"version": "1.0.0",
		})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}
