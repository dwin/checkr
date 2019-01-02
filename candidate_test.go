package checkr

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCandidate(t *testing.T) {
	// Load Key & Create Client
	apiKey := LoadKey(t)
	c := NewClient(apiKey)

	// Test Candidate w/Missing Requirement
	candidate := Candidate{
		FirstName:    randomdata.FirstName(randomdata.RandomGender),
		NoMiddleName: true,
		LastName:     randomdata.LastName(),
		DOB:          "1970-01-25",
		SSN:          "111-11-2001",
		Zipcode:      "49503",
	}
	_, err := c.CreateCandidate(candidate)
	require.Error(t, err, "Create should fail with missing email")

	// Test Candidate
	candidate = Candidate{
		FirstName:    randomdata.FirstName(randomdata.RandomGender),
		NoMiddleName: true,
		LastName:     randomdata.LastName(),
		Email:        randomdata.Email(),
		DOB:          "1970-01-25",
		SSN:          "111-11-2001",
		Zipcode:      "49503",
		CustomID:     randomdata.Digits(14),
	}
	createRes, err := c.CreateCandidate(candidate)
	require.NoError(t, err)
	assert.NotEmpty(t, createRes.ID)
	assert.Equal(t, candidate.FirstName, createRes.FirstName)

	// Retrieve New Candidate
	retRes, err := c.RetrieveCandidate(createRes.ID)
	require.NoError(t, err)
	assert.Equal(t, createRes.ID, retRes.ID)
	assert.Equal(t, createRes.FirstName, retRes.FirstName)

	// Update Candidate
	updateCan := *retRes
	updateCan.Email = randomdata.Email()
	updateCan.SSN = ""
	updateRes, err := c.UpdateCandidate(updateCan)
	require.NoError(t, err)
	assert.Equal(t, updateCan.ID, updateRes.ID)

	// List Candidates
	candidates, err := c.ListExistingCandidates()
	require.NoError(t, err)
	assert.NotNil(t, candidates)

}
