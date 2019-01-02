package checkr

import (
	"fmt"
	"net/http"
	"time"
)

// Document ...
// https://docs.checkr.com/#document
type Document struct {
	ID          string    `json:"id"`
	Object      string    `json:"object"`
	CreatedAt   time.Time `json:"created_at"`
	DownloadURI string    `json:"download_uri"`
	Filesize    int       `json:"filesize"`
	Filename    string    `json:"filename"`
	Type        string    `json:"type"`
	ContentType string    `json:"content_type"`
}

// Documents ...
type Documents struct {
	Document []Document `json:"data"`
	Object   string     `json:"object"`
	Count    int        `json:"count"`
}

// RetrieveDocument ...
func (c *Client) RetrieveDocument(documentID string) (*Document, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Document{}).SetError(&ErrorResponse{}).Get("/documents/" + documentID)
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Document), nil
}

// RetrieveCandidateDocuments ...
func (c *Client) RetrieveCandidateDocuments(candidateID string) (*Documents, error) {
	// Handle Request
	resp, err := c.R().SetResult(&Documents{}).SetError(&ErrorResponse{}).Get("/candidates/" + candidateID + "/documents/")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Documents), nil
}

// UploadCandidateDocumentFile ...
func (c *Client) UploadCandidateDocumentFile(candidateID, documentType, filepath string) (*Document, error) {
	// f, err := os.Open(filepath)
	// if err != nil {
	// 	return nil, err
	// }
	// Handle Request
	resp, err := c.R().SetQueryString("type="+documentType).SetFile("file", filepath).SetResult(&Document{}).SetError(&ErrorResponse{}).Post("/candidates/" + candidateID + "/documents")
	if err != nil {
		return nil, err
	}
	// Check for expected response
	if resp.StatusCode() != http.StatusCreated {
		errResp := resp.Error().(*ErrorResponse)
		err = fmt.Errorf("Checkr Error: %s", errResp.Error)
		return nil, err
	}

	return resp.Result().(*Document), nil
}
