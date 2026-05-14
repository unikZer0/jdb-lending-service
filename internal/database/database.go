package database

import (
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() error {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "apiservice")
	password := getEnv("DB_PASSWORD", "apiservice123")
	dbname := getEnv("DB_NAME", "apiservice_db")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	fmt.Println("Successfully connected to database!")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
