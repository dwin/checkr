package checkr

import (
	"testing"

	randomdata "github.com/pallinder/go-randomdata"
	"github.com/stretchr/testify/require"
)

func TestInvitation(t *testing.T) {
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
	require.NotNil(t, cand)

	// Create Invitation
	inv, err := c.CreateInvitation("driver_pro", cand.ID)
	require.NoError(t, err)
	require.NotNil(t, inv)
}
