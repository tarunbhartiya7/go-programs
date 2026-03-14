# Programs

Go programs that explain and demonstrate different features of the Go language.

## Prerequisites

- Go 1.25.0 or later

## Project Structure

Each program demonstrates specific Go features:

| Path | Features |
|------|----------|
| Root (`helloworld.go`, `arrays.go`, `slice.go`, `maps.go`, `pointers-exercise.go`, `functions.go`, `structures.go`, `errors.go`, `ranges.go`, `operators.go`) | Basics, arrays, slices, maps, pointers, functions, structs, errors, iteration |
| `bank/` | File I/O, switch, loops, defer |
| `notes-app/` | Structs, methods, JSON marshaling, packages |
| `struct-classes/` | Structs, methods, method receivers, HTTP client |
| `email/` | Packages, regex, unit testing |
| `goroutines/` | Goroutines and concurrency |
| `investment-calculator/` | Variables, input, arithmetic |
| `profit-calculator/` | Functions, calculations |
| `queue/` | Data structures |
| `pkg/display/` | Custom packages |
| `problems/` | Error handling, method receivers |

## Running Programs

**Root-level files:**
```bash
go run <file>.go
```
Example: `go run helloworld.go`

**Apps with their own modules:**
```bash
cd <folder>
go run .
```
Example: `cd bank && go run .`

## Important Go Commands

| Command | Description |
|---------|-------------|
| `go version` | Check current Go version |
| `go mod init example.com/app` | Create a new Go module |
| `go run <filename>` | Run a Go program |
| `goenv` | Manage multiple Go versions (like pyenv) |
| `goenv versions` | List available Go versions |
| `air` | Live reload for Go (like nodemon) |
