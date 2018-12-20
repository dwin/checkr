package checkr

import (
	"fmt"
	"net/http"
	"time"
)

// Report ...
type Report struct {
	ID                       string     `json:"id"`
	Object                   string     `json:"object"`
	URI                      string     `json:"uri"`
	Status                   string     `json:"status"`
	CreatedAt                time.Time  `json:"created_at"`
	CompletedAt              time.Time  `json:"completed_at"`
	RevisedAt                *time.Time `json:"revised_at"`
	TurnaroundTime           int        `json:"turnaround_time"`
	DueTime                  time.Time  `json:"due_time"`
	Adjudication             string     `json:"adjudication"`
	Package                  string     `json:"package"`
	Source                   string     `json:"source"`
	CandidateID              string     `json:"candidate_id"`
	SsnTraceID               string     `json:"ssn_trace_id"`
	SexOffenderSearchID      string     `json:"sex_offender_search_id"`
	NationalCriminalSearchID string     `json:"national_criminal_search_id"`
	CountyCriminalSearchIds  []string   `json:"county_criminal_search_ids"`
	MotorVehicleReportID     string     `json:"motor_vehicle_report_id"`
	StateCriminalSearchIds   []string   `json:"state_criminal_search_ids"`
	DocumentIds              []string   `json:"document_ids"`
	GeoIds                   []string   `json:"geo_ids"`
	ProgramID                string     `json:"program_id"`
}

// CreateReport ...
// ex. c.CreateReport("driver_pro","e44aa283528e6fde7d542194")
func (c *Client) CreateReport(pkg, candidateID string) (*Report, error) {
	body := map[string]string{
		"package":      pkg,
		"candidate_id": candidateID,
	}
	// Handle Request
	resp, err := c.R().SetBody(body).SetResult(&Report{}).SetError(&ErrorResponse{}).Post("/reports")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Report), nil
}

// RetrieveReport ...
func (c *Client) RetrieveReport(id string) (*Report, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Report{}).SetError(&ErrorResponse{}).Get("/reports/" + id)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Report), nil
}
