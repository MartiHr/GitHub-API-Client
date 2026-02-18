# GitHub API Client

Implements GItHub API HTTP client that will:

Read a text file given as command line argument to the program and parse different Github usernames – each username on separate line in the file.
Fetch GitHub users data in JSON format using public GitHub API: https://api.github.com/users/${username}
Fetch GitHub user repositories data in JSON format from: https://api.github.com/users/${username}/repos
Fetch information about programming languages in each repo from: https://api.github.com/repos/${username}/${repo-name}/languages
Parse the JSON data using json.Unmarshal into appropriate data structures in Go (you could define only fields that are interesting, all fields exported = starting with capital letter).
Print a statistics report containing the information about the user, the number of user repositories, the distribution of programming languages according to their usage numbers (third URL), the total number of followers, number of forks for all repositories, by year distribution of user activity calculated using repositories creation and last update dates. Format the report and print it to the console as structured tables
