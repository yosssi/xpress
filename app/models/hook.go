package models

// Hook represents a GitHub hook.
type Hook struct {
	Ref        string     `json:"ref"`
	HeadCommit Commit     `json:"head_commit"`
	Repository Repository `json:"repository"`
}
