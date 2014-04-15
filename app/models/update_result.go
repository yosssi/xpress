package models

// An UpdateResult represents an update result.
type UpdateResult struct {
	ID      string `json:"_id"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Version int64  `json:"_version"`
}
