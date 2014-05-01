package models

import "time"

// An Article represents an article.
type Article struct {
	ID          string
	FileName    string
	Title       string
	PublishedAt time.Time
	UpdatedAt   time.Time
}

// NewArticle generates an Article and returns it.
func NewArticle(id string, fileName string, title string, publishedAt time.Time, updatedAt time.Time) *Article {
	return &Article{ID: id, FileName: fileName, Title: title, PublishedAt: publishedAt, UpdatedAt: updatedAt}
}

// NewArticleFromMarkdown parses a markdown, generates an Article and returns it.
func NewArticleFromMarkdown(markdown string) {

}
