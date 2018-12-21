package checkr

import (
	"fmt"
	"net/http"
)

// AdverseItem ...
type AdverseItem struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Text   string `json:"text"`
}

// AdverseItems ...
type AdverseItems struct {
	Data   []AdverseItem `json:"data"`
	Object string        `json:"object"`
	Count  int           `json:"count"`
}

// ListAdverseItems ...
func (c *Client) ListAdverseItems(reportID string) (*[]AdverseItem, error) {
	// Handle Request
	resp, err := c.R().SetResult(&AdverseItems{}).SetError(&ErrorResponse{}).Get("/reports/" + reportID + "/adverse_items")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	resData := resp.Result().(*AdverseItems)

	return &resData.Data, nil
}
