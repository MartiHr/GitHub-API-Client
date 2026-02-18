package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadUserNamesFromFile(fileName string) ([]string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	var usernames []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		username := scanner.Text()

		// Normalize username
		username = strings.TrimSpace(username)

		if username != "" {
			usernames = append(usernames, scanner.Text())
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}
