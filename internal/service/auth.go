package service

import (
	"apiservice/internal/models"
	"apiservice/internal/provider"
	"apiservice/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req models.RegisterRequest) error
	Login(req models.LoginRequest) (*models.Auth, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) Register(req models.RegisterRequest) error {
	exists, err := s.authRepo.EmailExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already registered")
	}

	exists, err = s.authRepo.UsernameExists(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.authRepo.CreateUser(req.Username, req.Email, string(hashedPassword))
}

func (s *authService) Login(req models.LoginRequest) (*models.Auth, error) {
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := provider.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &models.Auth{
		Token: token,
	}, nil
}
