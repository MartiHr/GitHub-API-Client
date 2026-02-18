package models

// Language represents the languages used in a single GitHub repository.
type Language struct {
	RepoName  string         `json:"repo_name"`
	LangStats map[string]int `json:"languages"`
}

// Languages is a slice of Language entries.
type Languages []Language
