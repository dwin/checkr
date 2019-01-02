package checkr

import "time"

// Invitation ...
type Invitation struct {
	ID            string     `json:"id"`
	Status        string     `json:"status"`
	URI           string     `json:"uri"`
	InvitationURL string     `json:"invitation_url"`
	CompletedAt   *time.Time `json:"completed_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	ExpiresAt     time.Time  `json:"expires_at"`
	Package       string     `json:"package"`
	Object        string     `json:"object"`
	CreatedAt     time.Time  `json:"created_at"`
	CandidateID   string     `json:"candidate_id"`
}
