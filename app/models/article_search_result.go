package models

// An ArticleSearchResult represents a search result of users.
type ArticleSearchResult struct {
	Hits struct {
		Hits []struct {
			ID     string        `json:"_id"`
			Index  string        `json:"_index"`
			Score  float64       `json:"_score"`
			Source ArticleSource `json:"_source"`
			Type   string        `json:"_type"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    int64   `json:"total"`
	} `json:"hits"`
	SearchResult
}

// Article generates an article and returns it.
func (u *ArticleSearchResult) Article() *Article {
	hits := u.Hits.Hits

	if hits == nil || len(hits) == 0 {
		return nil
	}

	hit := hits[0]
	source := hit.Source

	return NewArticle(hit.ID, source.FileName, source.Title, source.PublishedAt, source.UpdatedAt)
}
