package models

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	UserViewType      string    `json:"user_view_type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              any       `json:"name"`
	Company           any       `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             any       `json:"email"`
	Hireable          any       `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   any       `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// String method for individual User
func (u *User) String() string {
	return fmt.Sprintf("User<%d> - Login: %s, Name: %s", u.ID, u.Login, u.Name)
}

type Users []User

func (u *Users) String() string {
	if len(*u) == 0 {
		return "No repositories"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Repositories (%d):\n", len(*u)))
	sb.WriteString(strings.Repeat("=", 50) + "\n")

	for i, user := range *u {
		if i > 0 {
			sb.WriteString("\n" + strings.Repeat("-", 50) + "\n")
		}
		sb.WriteString(user.String())
	}

	return sb.String()
}
