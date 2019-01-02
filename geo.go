package checkr

import "time"

// Geo ...
// https://docs.checkr.com/#geo
type Geo struct {
	ID        string     `json:"id"`
	Object    string     `json:"object"`
	URI       string     `json:"uri"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `json:"name"`
	City      string     `json:"city"`
	State     string     `json:"state"`
	DeletedAt *time.Time `json:"deleted_at"`
}
