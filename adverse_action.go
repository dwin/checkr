package checkr

import (
	"fmt"
	"net/http"
	"time"
)

// AdverseAction ...
type AdverseAction struct {
	ID                    string        `json:"id"`
	Object                string        `json:"object"`
	URI                   string        `json:"uri"`
	CreatedAt             *time.Time    `json:"created_at"`
	Status                string        `json:"status"`
	ReportID              string        `json:"report_id"`
	PostNoticeScheduledAt *time.Time    `json:"post_notice_scheduled_at"`
	PostNoticeReadyAt     *time.Time    `json:"post_notice_ready_at"`
	CanceledAt            *time.Time    `json:"canceled_at"`
	AdverseItems          []AdverseItem `json:"adverse_items"`
}

// CreateAdverseAction ...
func (c *Client) CreateAdverseAction(reportID string, adverseItemIDs []string, postNoticeScheduledAt ...time.Time) (*AdverseAction, error) {
	body := map[string]interface{}{
		"adverse_item_ids": adverseItemIDs,
	}
	if len(postNoticeScheduledAt) > 0 {
		body["post_notice_scheduled_at"] = postNoticeScheduledAt[0].Format(time.RFC3339)
	}
	// Handle Request
	resp, err := c.R().SetBody(body).SetResult(&AdverseAction{}).SetError(&ErrorResponse{}).Post("/reports/" + reportID + "/adverse_actions")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*AdverseAction), nil
}
