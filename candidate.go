package checkr

import (
	"fmt"
	"net/http"
	"time"
)

// Candidate represents a candidate to be screened.
// https://docs.checkr.com/#candidate
type Candidate struct {
	ID                          string     `json:"id,omitempty"`
	Object                      string     `json:"object,omitempty"`
	URI                         string     `json:"uri,omitempty"`
	CreatedAt                   *time.Time `json:"created_at,omitempty"`
	FirstName                   string     `json:"first_name,omitempty"`
	MiddleName                  string     `json:"middle_name,omitempty"`
	NoMiddleName                bool       `json:"no_middle_name,omitempty"`
	LastName                    string     `json:"last_name,omitempty"`
	MotherMaidenName            string     `json:"mother_maiden_name,omitempty"`
	Email                       string     `json:"email,omitempty"`
	Phone                       string     `json:"phone,omitempty"`
	Zipcode                     string     `json:"zipcode,omitempty"`
	DOB                         string     `json:"dob,omitempty"`
	SSN                         string     `json:"ssn,omitempty"`
	DriverLicenseNumber         string     `json:"driver_license_number,omitempty"`
	DriverLicenseState          string     `json:"driver_license_state,omitempty"`
	PreviousDriverLicenseNumber string     `json:"previous_driver_license_number,omitempty"`
	PreviousDriverLicenseState  string     `json:"previous_driver_license_state,omitempty"`
	CopyRequested               bool       `json:"copy_requested,omitempty"`
	CustomID                    string     `json:"custom_id,omitempty"`
	ReportIDs                   []string   `json:"report_ids,omitempty"`
	GeoIDs                      []string   `json:"geo_ids,omitempty"`
}

// CreateCandidate ...
func (c *Client) CreateCandidate(candidate Candidate) (*Candidate, error) {
	// Handle Request
	resp, err := c.R().SetBody(candidate).SetResult(&Candidate{}).SetError(&ErrorResponse{}).Post("/candidates")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusCreated {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Candidate), nil
}

type candidatesResp struct {
	Data []Candidate `json:"data"`
}

// RetrieveCandidate ...
func (c *Client) RetrieveCandidate(id string) (*Candidate, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Candidate{}).SetError(&ErrorResponse{}).Get("/candidates/" + id)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Candidate), nil
}
