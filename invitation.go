package checkr

import (
	"fmt"
	"net/http"
	"time"
)

// Invitation ...
// Represents a background check invitation. The candidate will receive an email to submit their information.
// https://docs.checkr.com/#invitation
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

// CreateInvitation is used to create a new Invitation.
func (c *Client) CreateInvitation(pkg, candidateID string) (*Invitation, error) {
	invite := Invitation{
		Package:     pkg,
		CandidateID: candidateID,
	}
	// Handle Request
	resp, err := c.R().SetBody(invite).SetResult(&Invitation{}).SetError(&ErrorResponse{}).Post("/invitations")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusCreated {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Invitation), nil
}

// CancelInvitation ...
func (c *Client) CancelInvitation(invitationID string) (*Invitation, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Invitation{}).SetError(&ErrorResponse{}).Delete("/invitations/" + invitationID)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Invitation), nil
}

// RetrieveInvitation ...
func (c *Client) RetrieveInvitation(invitationID string) (*Invitation, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Invitation{}).SetError(&ErrorResponse{}).Get("/invitations/" + invitationID)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Invitation), nil
}

// RetrieveInvitations ...
// status		string 	values: 'pending', 'completed', 'expired'
// candidate_id	string
// ex. params := map[string]string{
// 	"status":"pending",
// }
// inv, err := c.RetrieveInvitations(params)
func (c *Client) RetrieveInvitations(params ...map[string]string) (*Invitation, error) {
	// Handle Request
	req := c.R().SetResult(&Invitation{}).SetError(&ErrorResponse{})
	if len(params) > 1 {
		req.SetQueryParams(params[0])
	}
	resp, err := req.Get("/invitations")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Invitation), nil
}
