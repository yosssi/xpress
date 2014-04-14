package models

// A User represents a user.
type User struct {
	ID          string
	AccessToken string
	GitHubID    float64
}

// NewUser generates a user and returns it.
func NewUser(id string, accessToken string, gitHubID float64) *User {
	return &User{ID: id, AccessToken: accessToken, GitHubID: gitHubID}
}
