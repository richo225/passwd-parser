<h1 align="center">Passwd Parser</h1>

<p align="center"> Tiny CLI for parsing a password file and outputting to CSV/JSON
    <br> 
</p>

## 📝 Table of Contents

- [Getting Started](#getting_started)
- [Usage](#usage)

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Make sure you have Go version 1.19 or later.

```
$ go version
go version go1.19 darwin/arm64
```

### Installing

Run the executable:

```
$ bin/passwd-parser
```

If you want to use your own password file, upload it in the lib directory and either re-build the binary or compile and execute with go run:

```
$ mv /path-to-your-file /lib/passwd

$ go build -o bin/passwd-parser
// OR
$ go run main.go
```

## 🎈 Usage <a name="usage"></a>

The program parses the `lib/passwd` file. You can specify with flags the `-path` of the output file and the `-format` that you require(default is json)

```
-format string
    output format for the user information eg, csv, json (default "json")
-path string
    path to export file
```

### Examples

When no path is provided, user info will be written to STDOUT

<img width="336" alt="passwd-parser-csv" src="https://user-images.githubusercontent.com/18379191/184295952-4c96cb12-b95d-4681-89d2-803fd9d2f948.png">

---

When a path is provided, user info will be written to the existing file or created

<img width="385" alt="parser-json-file" src="https://user-images.githubusercontent.com/18379191/184295829-40f8545c-630e-44af-8959-bcfb94534b16.png">








