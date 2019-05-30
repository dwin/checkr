package checkr

import (
	"testing"

	randomdata "github.com/pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {
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

	// Get Candidate Documents
	docs, err := c.RetrieveCandidateDocuments(cand.ID)
	require.NoError(t, err)
	assert.Empty(t, docs.Document, "Documents should not be available for newly created candidate")

	// Upload Candidate Document
	newDoc, err := c.UploadCandidateDocumentFile(cand.ID, "driver_license", "testimg.png")
	require.NoError(t, err)
	assert.NotEmpty(t, newDoc.ID)

	// Get Document
	doc, err := c.RetrieveDocument(newDoc.ID)
	require.NoError(t, err)
	assert.Equal(t, newDoc.ID, doc.ID)
}
