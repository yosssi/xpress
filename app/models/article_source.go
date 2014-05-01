package models

import "time"

// An ArticleSource represents an article source.
type ArticleSource struct {
	FileName    string    `json:"file_name"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:published_at`
	UpdatedAt   time.Time `json:updated_at`
}
