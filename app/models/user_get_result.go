package models

// UserGetResult represents a get result of a user.
type UserGetResult struct {
	Found   bool       `json:"found"`
	ID      string     `json:"_id"`
	Index   string     `json:"_index"`
	Source  UserSource `json:"_source"`
	Type    string     `json:"_type"`
	Version int64      `json:"_version"`
}

// User generates a user and return it.
func (u *UserGetResult) User() *User {
	return NewUser(u.ID, u.Source.AccessToken, u.Source.GithubID)
}
