package checkr

import (
	"encoding/json"
	"time"
)

// WebhookRequest ...
type WebhookRequest struct {
	ID         string    `json:"id"`
	Object     string    `json:"object"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	WebhookURL string    `json:"webhook_url"`
	Data       struct {
		Object json.RawMessage `json:"object"`
	} `json:"data"`
}
