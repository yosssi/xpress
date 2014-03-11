package models

import (
	"os"

	"github.com/yosssi/xpress/app/consts"
)

// A GitHub represents information about GitHub.
type GitHub struct {
	ClientID     string
	ClientSecret string
}

// NewGitHub generates a GitHub and returns it.
func NewGitHub() *GitHub {
	return &GitHub{ClientID: os.Getenv(consts.EnvGitHubClientID), ClientSecret: os.Getenv(consts.EnvGitHubClientSecret)}
}
