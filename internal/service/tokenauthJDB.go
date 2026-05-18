package service

import (
	"apiservice/internal/models"
	"apiservice/internal/provider"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var tokenCache = provider.TokenCacheInstance

func GetJDBToken(url, userID, secretID, reqID string) (string, *models.JDBErrorResponse, error) {
	if token, valid := tokenCache.Get(userID); valid {
		fmt.Printf("cache hit na : %s\n", userID)
		return token, nil, nil
	}

	fmt.Printf("Fetching new token for user: %s\n", userID)

	fmt.Println("Fetching new token from JDB API")
	payload := models.AuthRequest{
		RequestID: reqID,
		UserID:    userID,
		SecretID:  secretID,
	}
	body, _ := json.Marshal(payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if err != nil {
		fmt.Println("here ??fr")
		return "", nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("here ??")
		fmt.Println("err:", err)
		return "", nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var jdbErr models.JDBErrorResponse
		json.Unmarshal(respData, &jdbErr)
		fmt.Println("Error from Server: ??", string(respData))
		return "", &jdbErr, nil
	}
	var res models.AuthResponse
	if err := json.Unmarshal(respData, &res); err != nil {
		return "", nil, err
	}
	tokenCache.Set(userID, res.Data.Token, 1*time.Hour)
	return res.Data.Token, nil, nil
}
