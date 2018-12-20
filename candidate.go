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

// Candidates ...
type Candidates struct {
	Data []Candidate
}

// CreateCandidate ...
/*
	Note: To support use of the Invitations API (in which case, the invitation apply form would collect a Candidate's personal information), the only strictly required Candidate attribute is email.
	However, in most cases, additional personal information will be required to request a Report for a Candidate.
	Required attributes would include:
	first_name
	middle_name or no_middle_name
	last_name
	dob
	If the Report's package includes any type of criminal check, the following will be required:
	ssn
	zipcode
	If the Report's package includes a Motor Vehicle Report (MVR), the following will be required:
	driver_license_number
	driver_license_state
	Validation for these attributes is performed when requesting a Report, as the requirements depend on the Package.
*/
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

// UpdateCandidate ...
/*
	Note: non-null fields cannot be updated after a Report has been ordered for a Candidate with the exception of the following fields:
	email
	previous_driver_license_number
	previous_driver_license_state
	copy_requested
	custom_id
	geo_ids
	Updating geo_ids will replace all existing geos. If you want to keep existing geos, include their ids in the update request.
*/
func (c *Client) UpdateCandidate(candidate Candidate) (*Candidate, error) {
	// Handle Request
	resp, err := c.R().SetBody(candidate).SetResult(&Candidate{}).SetError(&ErrorResponse{}).Post("/candidates/" + candidate.ID)
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

// ListExistingCandidates ...
// Valid Query Parameters:
// email:	string
// full_name:	string
// adjudication:	string
// custom_id:	string
// created_after:	date
// format: YYYY-MM-DD
// filter candidates created after this timestamp
// created_before:	date
// format: YYYY-MM-DD
// filter candidates created before this timestamp
// geo_id:	string
// per_page:	integer
// between 0 and 100
// page:	integer
// greater than or equal to 1
//
// example:
//  params := map[string]string{
// 	   "email":"test@tld.com",
//  }
func (c *Client) ListExistingCandidates(params ...map[string]string) (*Candidates, error) {
	// Handle Query Params
	req := c.R().SetResult(&Candidates{}).SetError(&ErrorResponse{})
	if len(params) > 1 {
		req.SetQueryParams(params[0])
	}
	// Handle Request
	resp, err := req.Get("/candidates")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Candidates), nil
}
