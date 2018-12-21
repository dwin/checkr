package checkr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func LoadKey(t *testing.T) string {
	key := os.Getenv("CHECKR_KEY")
	require.NotEmpty(t, key, "Env var 'CHECKR_KEY' must be set for testing")
	return key
}

func TestClient(t *testing.T) {
	// Test User-set API URL
	url := "http://localhost"
	c := NewClient("key", url)
	assert.Equal(t, url, c.HostURL, "User given URL should be set on client")
}
