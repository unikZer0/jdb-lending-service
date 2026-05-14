package models

// JDB API Types
type AuthRequest struct {
	RequestID string `json:"requestId"`
	UserID    string `json:"userId"`
	SecretID  string `json:"secretId"`
}

type AuthResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

type JDBErrorResponse struct {
	Error struct {
		Code      int    `json:"code"`
		Status    string `json:"status"`
		Resource  string `json:"resource"`
		Reason    string `json:"reason"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
	} `json:"error"`
}

type LendingRequest struct {
	RequestID string `json:"requestId"`
	CIF       string `json:"cif"`
	Language  string `json:"language"`
}
