package checkr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func LoadKey(t *testing.T) string {
	key := os.Getenv("CHECKR_KEY")
	require.NotEmpty(t, key, "Env var 'CHECKR_KEY' must be set for testing")
	return key
}
