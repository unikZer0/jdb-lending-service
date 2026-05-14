package service

import (
	"apiservice/internal/models"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLendingURL(url, token, reqID, cif, language, secretKey string) (map[string]interface{}, *models.JDBErrorResponse, error) {
	payload := models.LendingRequest{
		RequestID: reqID,
		CIF:       cif,
		Language:  "EN",
	}

	body, _ := json.Marshal(payload)
	signature := ComputeHmac256(string(body), secretKey)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("SignedHash", signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error from Lending Server:", string(respData))
		var jdbErr models.JDBErrorResponse
		json.Unmarshal(respData, &jdbErr)
		return nil, &jdbErr, nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, nil, err
	}
	fmt.Println(result)
	return result, nil, nil
}
func ComputeHmac256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
