package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Home  string `json:"home"`
	Shell string `json:"shell"`
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(collectUsers())
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func collectUsers() (users []User) {
	f, err := os.Open("hr/passwd")
	handleError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ':'

	lines, err := reader.ReadAll()
	handleError(err)

	for _, line := range lines {
		id, err := strconv.ParseInt(line[2], 10, 64)
		handleError(err)

		// Don't want system users here
		if id < 100 {
			continue
		}

		user := User{
			Id:    int(id),
			Name:  line[0],
			Home:  line[5],
			Shell: line[6],
		}

		users = append(users, user)
	}

	return
}
