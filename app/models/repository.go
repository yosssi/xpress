package models

// Repository represents a GitHub repository.
type Repository struct {
	ID           int64  `json:"id"`
	MasterBranch string `json:"master_branch"`
	Name         string `json:"name"`
	URL          string `json:"url"`
}
