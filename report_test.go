package checkr

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	// Load Key & Create Client
	apiKey := LoadKey(t)
	c := NewClient(apiKey)

	// Create Candidate for Testing
	candidate := Candidate{
		FirstName:           randomdata.FirstName(randomdata.RandomGender),
		NoMiddleName:        true,
		LastName:            randomdata.LastName(),
		Email:               randomdata.Email(),
		DOB:                 "1970-01-25",
		SSN:                 "111-11-2001",
		Zipcode:             "49503",
		DriverLicenseNumber: "F1112001",
		DriverLicenseState:  "CA",
		CustomID:            randomdata.Digits(14),
	}
	cand, err := c.CreateCandidate(candidate)
	require.NoError(t, err)
	// Create Report
	report, err := c.CreateReport("driver_pro", cand.ID)
	require.NoError(t, err)
	// Update Report
	report, err = c.UpdateReport(report.ID, "driver_pro", "")
	require.NoError(t, err)
	// Retrieve Report
	report, err = c.RetrieveReport(report.ID)
	require.NoError(t, err)
	assert.NotEmpty(t, report.ID)
}
