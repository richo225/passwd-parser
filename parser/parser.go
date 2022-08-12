package parser

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

const PASSWD_FILE_PATH = "lib/passwd"

func ParseFlags() (path, format string) {
	flag.StringVar(&path, "path", "", "path to export file")
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

func CollectUsers() (users []User) {
	f, err := os.Open(PASSWD_FILE_PATH)
	HandleError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ':'

	lines, err := reader.ReadAll()
	HandleError(err)

	for _, line := range lines {
		id, err := strconv.ParseInt(line[2], 10, 64)
		HandleError(err)

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

func WriteOutput(path, format string, users []User) {
	var output io.Writer
	if path != "" {
		// write user info to path in specified format
		f, err := os.Create(path)
		HandleError(err)
		defer f.Close()
		output = f
	} else {
		// write user info to stdout
		output = os.Stdout
	}

	if format == "csv" {
		output.Write([]byte("id,name,home,shell\n"))
		writer := csv.NewWriter(output)
		for _, user := range users {
			err := writer.Write([]string{strconv.Itoa(user.Id), user.Name, user.Home, user.Shell})
			HandleError(err)
		}
		writer.Flush()
	}

	if format == "json" {
		data, err := json.MarshalIndent(users, "", "  ")
		HandleError(err)
		output.Write(data)
	}
}

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
