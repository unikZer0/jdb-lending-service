package repository

import (
	"apiservice/internal/database"
	"apiservice/internal/models"
)

type AuthRepository interface {
	EmailExists(email string) (bool, error)
	UsernameExists(username string) (bool, error)
	CreateUser(username, email, hashedPassword string) error
	GetUserByEmail(email string) (*models.User, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) EmailExists(email string) (bool, error) {
	var exists bool
	err := database.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email)
	return exists, err
}

func (r *authRepository) UsernameExists(username string) (bool, error) {
	var exists bool
	err := database.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username)
	return exists, err
}

func (r *authRepository) CreateUser(username, email, hashedPassword string) error {
	var userID int
	err := database.DB.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		username, email, hashedPassword,
	).Scan(&userID)
	return err
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
