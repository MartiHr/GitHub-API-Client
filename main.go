package main

import (
	"Homework1/fetchers"
	"Homework1/models"
	"Homework1/parsers"
	"Homework1/visualisation"
	"fmt"
	"os"
)

const BaseUsersTemplate = "https://api.github.com/users/%s"
const BaseReposTemplate = "https://api.github.com/users/%s/repos"
const BaseLanguagesTemplate = "https://api.github.com/repos/%s/languages"

func main() {
	fmt.Print("Enter a file name: ")

	var fileName string
	n, err := fmt.Scanln(&fileName)
	if n != 1 || err != nil {
		fmt.Println("Error reading file name")
		os.Exit(1)
	}

	usernames, err := parsers.ReadUserNamesFromFile(fileName)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	var users models.Users
	var usersRepos []models.Repositories
	var languages models.Languages

	for _, username := range usernames {

		fmt.Println(username)

		// USERS
		//https://api.github.com/users/${username}
		userUrl := fmt.Sprintf(BaseUsersTemplate, username)
		fmt.Println(userUrl)

		user, err := FetchAndParse[models.User](userUrl)

		if err != nil {
			fmt.Println("Error getting user")
			os.Exit(1)
		}

		users = append(users, *user)

		fmt.Println()

		// REPOS
		// https://api.github.com/users/${username}/repos
		reposUrl := fmt.Sprintf(BaseReposTemplate, username)
		fmt.Println(reposUrl)

		repos, err := FetchAndParse[models.Repositories](reposUrl)

		if err != nil {
			fmt.Println("Error getting repos")
			os.Exit(1)
		}

		usersRepos = append(usersRepos, *repos)

		// Languages
		for _, repo := range *repos {
			repoFullName := repo.FullName
			langURL := fmt.Sprintf(BaseLanguagesTemplate, repoFullName)
			fmt.Printf("Fetching languages for %s...\n", repoFullName)

			langMap, err := FetchAndParse[map[string]int](langURL)
			if err != nil {
				fmt.Printf("Error getting languages for %s: %v\n", repoFullName, err)
				continue
			}

			// Store in languages list
			languages = append(languages, models.Language{
				RepoName:  repoFullName,
				LangStats: *langMap,
			})
		}

	}

	// (Optional) Print results with default Stringer implementation
	//fmt.Println(users)
	//fmt.Println(usersRepos)
	//fmt.Println(languages)

	visualisation.VisualizeStats(users, usersRepos, languages)
}

// FetchAndParse generic function to fetch and parse JSON data
func FetchAndParse[T any](url string) (*T, error) {
	respData, err := fetchers.FetchJSON(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching %w", err)
	}

	parsed, err := parsers.ParseJSON[T](respData)

	if err != nil {
		return nil, fmt.Errorf("error parsing %w", err)
	}

	return parsed, nil
}
