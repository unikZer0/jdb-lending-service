package handler

import (
	"apiservice/internal/models"
	"apiservice/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	_, err := repository.Register(req)
	if err != nil {
		if err.Error() == "email already registered" || err.Error() == "username already taken" {
			return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "regsiter successfully"})
}

func (h *Handler) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	response, err := repository.Login(req)
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
