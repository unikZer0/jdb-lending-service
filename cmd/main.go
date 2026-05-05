package main

import (
	"apiservice/internal/lending"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}
	lendingHandler := lending.NewHandler(
		os.Getenv("JDB_AUTH_URL"),
		os.Getenv("JDB_USER_ID"),
		os.Getenv("JDB_SECRET_ID"),
		os.Getenv("JDB_LENDING_URL"),
		os.Getenv("JDB_SECRET_KEY"),
	)
	e.POST("/request-lending", lendingHandler.HandleDigitalLending)
	e.GET("/", lending.HealthChecking)

	e.Logger.Fatal(e.Start(":8080"))
}
