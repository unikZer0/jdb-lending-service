package repository

import (
	"apiservice/internal/database"
	"apiservice/internal/models"
	"apiservice/internal/provider"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Register(req models.RegisterRequest) (string, error) {
	var exists bool
	err := database.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email)
	if err != nil {
		return "", fmt.Errorf("error checking email: %w", err)
	}
	if exists {
		return "", errors.New("email already registered")
	}

	err = database.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", req.Username)
	if err != nil {
		return "", fmt.Errorf("error checking username: %w", err)
	}
	if exists {
		return "", errors.New("username already taken")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	// Insert user
	var userID int
	err = database.DB.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		req.Username, req.Email, string(hashedPassword),
	).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	return "", nil
}

func Login(req models.LoginRequest) (*models.Auth, error) {
	// Get user by email
	var user models.User
	err := database.DB.Get(&user, "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = $1", req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := provider.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &models.Auth{
		Token: token,
	}, nil
}
