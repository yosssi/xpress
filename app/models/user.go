package models

// A User represents a user.
type User struct {
	ID          string
	AccessToken string
	GitHubID    int64
}

// NewUser generates a user and returns it.
func NewUser(id string, accessToken string, gitHubID int64) *User {
	return &User{ID: id, AccessToken: accessToken, GitHubID: gitHubID}
}
