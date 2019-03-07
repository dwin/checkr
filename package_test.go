package checkr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	randomdata "github.com/Pallinder/go-randomdata"
)

func TestPackage(t *testing.T) {
	// Load Key & Create Client
	apiKey := LoadKey(t)
	c := NewClient(apiKey)

	pkgName := randomdata.SillyName()
	screenings := Screening{
		Type: "ssn_trace",
	}

	// Create Package
	newPkg, err := c.CreatePackage(pkgName, pkgName+"_only", screenings)
	require.NoError(t, err)
	assert.NotEmpty(t, newPkg.ID, "New Package ID should not be empty")

	// List Packages
	pkgs, err := c.ListPackages()
	require.NoError(t, err)
	assert.NotEmpty(t, pkgs, "List of Packages should not be empty")

	// Retrieve Package
	retrieveID := pkgs.Packages[0].ID
	pkg, err := c.RetrievePackage(retrieveID)
	require.NoError(t, err)
	assert.Equal(t, retrieveID, pkg.ID)
}
