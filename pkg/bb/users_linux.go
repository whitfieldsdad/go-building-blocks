package bb

import (
	"bufio"
	"os"
	"strings"
)

func listUsers() ([]User, error) {
	usernames, err := listUsernames()
	if err != nil {
		return nil, err
	}
	var users []User
	for _, username := range usernames {
		user, err := GetUser(username)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func listUsernames() ([]string, error) {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var usernames []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		tokens := strings.Split(line, ":")
		if len(tokens) < 7 {
			continue
		}
		username := tokens[0]
		usernames = append(usernames, username)
	}
	return usernames, nil
}
