package main

import (
	"apiservice/internal/lending"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}
	e.POST("/request-lending", lending.HandleDigitalLending)
	e.GET("/", lending.HealthChecking)

	e.Logger.Fatal(e.Start(":8080"))
}
