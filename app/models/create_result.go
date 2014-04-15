package models

// A CreateResult represents a create result.
type CreateResult struct {
	Created bool   `json:"created"`
	ID      string `json:"_id"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Version int64  `json:"_version"`
}
