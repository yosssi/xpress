package models

// A SearchResult represents a search result.
type SearchResult struct {
	Error    string `json:"error"`
	Shards   Shards `json:"_shards"`
	Status   int16  `json:"status"`
	TimedOut bool   `json:"timed_out"`
	Took     int64  `json:"took"`
}
