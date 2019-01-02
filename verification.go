package checkr

import "time"

// Verification ...
type Verification struct {
	ID               string     `json:"id"`
	Object           string     `json:"object"`
	URI              string     `json:"uri"`
	CreatedAt        time.Time  `json:"created_at"`
	CompletedAt      *time.Time `json:"completed_at"`
	VerificationType string     `json:"verification_type"`
	VerificationURL  string     `json:"verification_url"`
}
