package auth

type AuthRequest struct {
	RequestID string `json:"requestId"`
	UserID    string `json:"userId"`
	SecretID  string `json:"secretId"`
}
type AuthResponse struct {
	TimeStamp     string `json:"timstamp"`
	Success       bool   `json:"success"`
	Message       string `json:"massage"`
	TransactionID string `json:"transactionId"`
	Data          struct {
		Token string `json:"token"`
	} `json:"data"`
}
type JDBErrorResponse struct {
	Timestamp     string          `json:"timestamp"`
	Success       bool            `json:"success"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	TransactionID string          `json:"transactionId"`
	NanoTime      int64           `json:"nanoTime"`
	DebugMessage  string          `json:"debugMessage"`
	Error         *JDBErrorDetail `json:"error,omitempty"`
}

type JDBErrorDetail struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type LendingRequest struct {
	RequestID string `json:"requestId"`
	CIF       string `json:"cif"`
	Language  string `json:"language"`
}
