package models

// Status -
type Status struct {
	Authorized bool `json:"authorized"`
	Version    int  `json:"version"`
}
