package models

// UserSource represents a user source.
type UserSource struct {
	AccessToken string `json:"access_token"`
	GithubID    int64  `json:"github_id"`
}
