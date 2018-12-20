package checkr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	// Load Key & Create Client
	apiKey := LoadKey(t)
	c := NewClient(apiKey)

	report, err := c.RetrieveReport("")
	require.NoError(t, err)
	fmt.Println(report)
}
