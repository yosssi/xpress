package models

// A Shards represents a shards.
type Shards struct {
	Failed     int64 `json:"failed"`
	Successful int64 `json:"successful"`
	Total      int64 `json:"total"`
}
