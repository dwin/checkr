package checkr

import (
	"fmt"
	"net/http"

	"gopkg.in/resty.v1"
)

const checkrAPIURL = "https://api.checkr.com/v1"
const checkrDateFormat = "2006-01-02"

// Client ...
type Client struct {
	*resty.Client
}

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewClient ...
func NewClient(apiKey string, apiURL ...string) *Client {
	// Create Resty Client
	r := resty.New()
	// Basic Auth with API Key as Username and no password
	// https://docs.checkr.com/#authentication
	r.SetBasicAuth(apiKey, "")
	// Use Default URL unless given API URL
	if len(apiURL) > 0 {
		r.SetHostURL(apiURL[0])
	} else {
		r.SetHostURL(checkrAPIURL)
	}

	return &Client{r}
}

// Status returns error is API status response is not as expected
func (c *Client) Status() error {
	// Handle Request
	resp, err := c.R().SetError(&ErrorResponse{}).Get("/status")
	if err != nil {
		return err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return err
	}

	return nil
}
