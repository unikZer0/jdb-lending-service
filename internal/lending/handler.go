package lending

import (
	"apiservice/internal/auth"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	AuthURL    string
	UserID     string
	SecretID   string
	LendingURL string
	SecretKey  string
}

func NewHandler(authURL, userID, secretID, lendingURL, secretKey string) *Handler {
	return &Handler{
		AuthURL:    authURL,
		UserID:     userID,
		SecretID:   secretID,
		LendingURL: lendingURL,
		SecretKey:  secretKey,
	}
}
func (h *Handler) HandleDigitalLending(c echo.Context) error {

	var req struct {
		RequestID string
		CIF       string
		Language  string
	}
	c.Bind(&req)
	token, jdbErr, err := auth.GetJDBToken(h.AuthURL, h.UserID, h.SecretID, fmt.Sprintf("%d", time.Now().Unix()))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Connection Error"})
	}
	if jdbErr != nil {
		fmt.Println("here")
		return c.JSON(http.StatusServiceUnavailable, jdbErr)
	}
	result, lendingJdbErr, err := FetchLendingURL(h.LendingURL, token, req.CIF, req.RequestID, req.Language, h.SecretKey)

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
