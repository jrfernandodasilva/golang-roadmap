
# Golang roadmap

[![License](https://img.shields.io/github/license/jrfernandodasilva/golang-roadmap.svg)](LICENSE)
[![Written in Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![GitHub release](https://img.shields.io/github/v/release/jrfernandodasilva/golang-roadmap.svg)](https://github.com/jrfernandodasilva/golang-roadmap/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Downloads](https://img.shields.io/github/downloads/jrfernandodasilva/golang-roadmap/total.svg?v1)](https://github.com/jrfernandodasilva/golang-roadmap/releases)
[![Contributors](https://img.shields.io/github/contributors/jrfernandodasilva/golang-roadmap.svg)](https://github.com/jrfernandodasilva/golang-roadmap/graphs/contributors)
[![Documentation](https://img.shields.io/badge/docs-latest-blue.svg)](https://github.com/jrfernandodasilva/golang-roadmap/wiki)
[![Last Update](https://img.shields.io/github/last-commit/jrfernandodasilva/golang-roadmap.svg)](https://github.com/jrfernandodasilva/golang-roadmap/commits/main)

### Init project module 
```bash
go mod init golang-roadmap
```

### Check modules dependencies
```bash
go mod tidy -v
```

### Init project module 
```bash
go mod init golang-roadmap
```

### Run
```bash
go run main.go
```

### Build with specific name
```bash
go build -o golang-roadmap main.go
```

### Build to specific OS
```bash
GOOS=linux GOARCH=amd64 go build -o golang-roadmap main.go
GOOS=windows GOARCH=amd64 go build -o golang-roadmap.exe main.go
GOOS=darwin GOARCH=amd64 go build -o golang-roadmap main.go
GOOS=android GOARCH=arm64 go build -o golang-roadmap main.go
```

### Format code file
```bash
go fmt main.go
```

### Run tests
```bash
go test -v
```

### Run tests with specific package
```bash
go test -v ./functions
```

### Run tests with specific package and file
```bash
go test -v ./math/int_operations_test.go
```

### Run tests with specific package and file with specific function
```bash
go test -v ./math/int_operations_test.go -run TestAddInt
```

### Project Structure
See [project structure](https://github.com/jrfernandodasilva/golang-roadmap/wiki/Project-Structure) suggestion in Wiki
