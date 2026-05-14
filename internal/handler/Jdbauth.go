package handler

import (
	"apiservice/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type JDBAuthHandler struct {
	AuthURL    string
	UserID     string
	SecretID   string
	LendingURL string
	SecretKey  string
}

func NewJDBAuthHandler(authURL, userID, secretID, lendingURL, secretKey string) *JDBAuthHandler {
	return &JDBAuthHandler{
		AuthURL:    authURL,
		UserID:     userID,
		SecretID:   secretID,
		LendingURL: lendingURL,
		SecretKey:  secretKey,
	}
}
func (h *JDBAuthHandler) HandleDigitalLending(c echo.Context) error {

	var req struct {
		RequestID string `json:"requestId"`
		CIF       string `json:"cif"`
		Language  string `json:"language"`
	}
	c.Bind(&req)
	token, jdbErr, err := service.GetJDBToken(h.AuthURL, h.UserID, h.SecretID, fmt.Sprintf("%d", time.Now().Unix()))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Connection Error"})
	}
	if jdbErr != nil {
		fmt.Println("here")
		return c.JSON(http.StatusServiceUnavailable, jdbErr)
	}
	result, lendingJdbErr, err := service.FetchLendingURL(h.LendingURL, token, req.RequestID, req.CIF, req.Language, h.SecretKey)

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, jdbErr)
	}
	if lendingJdbErr != nil {
		return c.JSON(http.StatusBadRequest, lendingJdbErr)
	}
	return c.JSON(http.StatusOK, result)
}
func HealthChecking(ctx echo.Context) error {
	fmt.Println("hi")
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "hi",
	})
}
