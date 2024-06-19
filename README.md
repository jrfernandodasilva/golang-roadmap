
[![license](https://img.shields.io/github/license/jrfernandodasilva/golang-roadmap.svg)](LICENSE)

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
