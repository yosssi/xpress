package models

// A UserSearchResult represents a search result of users.
type UserSearchResult struct {
	Hits struct {
		Hits []struct {
			ID     string     `json:"_id"`
			Index  string     `json:"_index"`
			Score  float64    `json:"_score"`
			Source UserSource `json:"_source"`
			Type   string     `json:"_type"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    int64   `json:"total"`
	} `json:"hits"`
	SearchResult
}

// User generates a user and returns it.
func (u *UserSearchResult) User() *User {
	hits := u.Hits.Hits

	if hits == nil || len(hits) == 0 {
		return nil
	}

	hit := hits[0]
	source := hit.Source

	return NewUser(hit.ID, source.AccessToken, source.GithubID)
}
