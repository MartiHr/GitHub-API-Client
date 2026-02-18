# GitHub API Client

A command-line utility written in **Go** that aggregates and visualizes GitHub user data. This tool processes batches of usernames to generate detailed statistics regarding repository activity, language distribution, and engagement metrics.

## Features

The client automates data collection by interacting with the following GitHub API endpoints:
* **User Profiles:** `GET /users/{username}`
* **Repository Lists:** `GET /users/{username}/repos`
* **Language Statistics:** `GET /repos/{username}/{repo}/languages`

### Core Functionality
* **Batch Input:** Reads a plain text file provided as a command-line argument (one username per line).
* **Efficient Parsing:** Uses `json.Unmarshal` to map API responses into optimized Go structs, selecting only relevant exported fields.
* **Data Aggregation:** Consolidates data from multiple endpoints to provide a holistic view of user activity.

---

## Technical Statistics

The program generates a structured console report including:

* **Profile Overview:** Username and total follower count.
* **Repository Metrics:** Total number of repositories and cumulative fork counts.
* **Language Distribution:** Usage numbers for programming languages across all repositories.
* **Activity Timeline:** Year-by-year distribution of activity based on repository creation and "last updated" timestamps.

---

## Usage

Ensure you have [Go](https://go.dev/) installed, then run the program by passing a text file containing GitHub usernames.

### 1. Prepare User List (`users.txt`)
```text
google
octocat
microsoft
