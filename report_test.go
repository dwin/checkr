package checkr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	// Load Key & Create Client
	apiKey := LoadKey(t)
	c := NewClient(apiKey)
	// Get Candidate for Report Retrieval
	cand, err := c.RetrieveCandidate("aa192591adb0bfd08ae1969d")
	require.NoError(t, err)

	report, err := c.RetrieveReport(cand.ReportIDs[0])
	require.NoError(t, err)
	assert.NotEmpty(t, report.ID)
}
