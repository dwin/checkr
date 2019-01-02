package checkr

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// CheckrPackage ...
// https://docs.checkr.com/#package
type CheckrPackage struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	URI        string      `json:"uri"`
	CreatedAt  time.Time   `json:"created_at"`
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	Price      int         `json:"price"`
	Screenings []Screening `json:"screenings"`
}

// CheckrPackages ...
type CheckrPackages struct {
	Package      []CheckrPackage `json:"data"`
	Object       string          `json:"object"`
	NextHref     interface{}     `json:"next_href"`
	PreviousHref interface{}     `json:"previous_href"`
	Count        int             `json:"count"`
}

// Screening ...
type Screening struct {
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
}

// CreatePackage - slug will to converted to all lowercase
func (c *Client) CreatePackage(name, slug string, screenings ...Screening) (*CheckrPackage, error) {
	// Input Validation
	for _, v := range screenings {
		switch v.Type {
		case "county_criminal_search", "county_civil_search", "employment_verification":
			if len(v.Subtype) < 1 {
				return nil, fmt.Errorf("Screening %v requires valid subtype", v.Type)
			}
		}
	}
	// Set Body
	body := map[string]interface{}{
		"name":       name,
		"slug":       strings.ToLower(slug),
		"screenings": screenings,
	}
	// Handle Request
	resp, err := c.R().SetBody(body).SetResult(&CheckrPackage{}).SetError(&ErrorResponse{}).Post("/packages")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusCreated {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*CheckrPackage), nil
}

// ListPackages ...
func (c *Client) ListPackages() (*CheckrPackages, error) {
	// Handle Request
	resp, err := c.R().SetResult(&CheckrPackages{}).SetError(&ErrorResponse{}).Get("/packages")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*CheckrPackages), nil
}

// RetrievePackage ...
func (c *Client) RetrievePackage(packageID string) (*CheckrPackage, error) {
	// Handle Request
	resp, err := c.R().SetResult(&CheckrPackage{}).SetError(&ErrorResponse{}).Get("/packages/" + packageID)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*CheckrPackage), nil
}
