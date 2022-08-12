package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"passwd-parser/parser"
	"strconv"
)

func main() {
	path, format := parser.ParseFlags()
	users := parser.CollectUsers()

	var output io.Writer
	if path != "" {
		// write user info to path in specified format
		f, err := os.Create(path)
		parser.HandleError(err)
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
			parser.HandleError(err)
		}
		writer.Flush()
	}

	if format == "json" {
		data, err := json.MarshalIndent(users, "", "  ")
		parser.HandleError(err)
		output.Write(data)
	}
}
