package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJDBToken(url, userID, secretID, reqID string) (string, *JDBErrorResponse, error) {
	payload := AuthRequest{
		RequestID: reqID,
		UserID:    userID,
		SecretID:  secretID,
	}
	body, _ := json.Marshal(payload)
	fmt.Println("payload :", payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println("here ?? fr")
		return "", nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("here ??")
		return "", nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var jdbErr JDBErrorResponse
		json.Unmarshal(respData, &jdbErr)
		fmt.Println("Error from Server: ??", string(respData))
		return "", &jdbErr, nil
	}
	var res AuthResponse
	if err := json.Unmarshal(respData, &res); err != nil {
		return "", nil, err
	}
	return res.Data.Token, nil, nil
}
