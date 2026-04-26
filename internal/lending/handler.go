package lending

import (
	"apiservice/internal/auth"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleDigitalLending(c echo.Context) error {

	authURL := os.Getenv("JDB_AUTH_URL")
	userID := os.Getenv("JDB_USER_ID")
	secretID := os.Getenv("JDB_SECRET_ID")
	lendingURL := os.Getenv("JDB_LENDING_URL")
	secretKey := os.Getenv("JDB_SECRET_KEY")

	var req struct {
		RequestID string
		CIF       string
		Language  string
	}
	c.Bind(&req)
	token, jdbErr, err := auth.GetJDBToken(authURL, userID, secretID, fmt.Sprintf("%d", time.Now().Unix()))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Connection Error"})
	}
	if jdbErr != nil {
		return c.JSON(http.StatusServiceUnavailable, jdbErr)
	}
	result, lendingJdbErr, err := FetchLendingURL(lendingURL, token, req.CIF, req.RequestID, req.Language, secretKey)

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
