package handler

import (
	"apiservice/internal/models"
	"apiservice/internal/repository"
	"apiservice/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	authService service.AuthService
}

func NewHandler() *Handler {
	authRepo := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepo)

	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Register(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	err := h.authService.Register(req)
	if err != nil {
		if err.Error() == "email already registered" || err.Error() == "username already taken" {
			return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "registered successfully"})
}

func (h *Handler) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	response, err := h.authService.Login(req)
	if err != nil {
		if err.Error() == "invalid email or password" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login"})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Me(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "login successfully"})
}
