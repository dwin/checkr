package checkr

import "time"

// Program ...
// https://docs.checkr.com/#program
type Program struct {
	ID         string     `json:"id"`
	Object     string     `json:"object"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	GeoIds     []string   `json:"geo_ids"`
	PackageIds []string   `json:"package_ids"`
}
