package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Home  string `json:"home"`
	Shell string `json:"shell"`
}

func main() {
	fmt.Println("Hello World")
	path, format := parseFlags()
	users := collectUsers()
}

func parseFlags() (path, format string) {
	flag.StringVar(&path, "path", "passwd", "path to export file")
	flag.StringVar(&format, "format", "json", "output format for the user information eg, csv, json")

	flag.Parse()

	// Check validity of format flag
	format = strings.ToLower(format)
	if !slices.Contains([]string{"csv", "json"}, format) {
		fmt.Println("Error: invalid format. Use 'json' or 'csv' instead.")
		flag.Usage()
		os.Exit(1)
	}

	return
}

func collectUsers() (users []User) {
	f, err := os.Open("passwd")
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

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
