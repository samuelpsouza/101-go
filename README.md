# 101 Go

A comprehensive learning repository for Go programming language, covering fundamentals and concurrent programming patterns.

## 📚 Learning Resources

### Books
- **[A Linguagem de Programação Go](https://www.amazon.com.br/Linguagem-Programa%C3%A7%C3%A3o-Go-Alan-Donovan/dp/8575225464/)** - The Go Programming Language (Portuguese Edition)
  - Authors: Alan Donovan & Brian Kernighan
  - Foundational reference for Go programming concepts

### Online Courses
- **[Concurrent Programming in Go](https://www.pluralsight.com/)** - Pluralsight Course by Mike VanSickle
  - Advanced patterns for concurrent and parallel programming
  - Goroutines, channels, and synchronization techniques

## 📁 Project Structure

### 1. **ALinguagemDeProgramacaoGo**
Exercises and examples from "The Go Programming Language" book (Chapter 1).

**Topics Covered:**
- `echo*.go` - Command-line argument handling
- `dup*.go` - File reading and map data structures
- `fetch*.go` - HTTP requests and network I/O
- `server*.go` - Basic HTTP server implementation
- `lissajous.go` - Graphics and mathematical visualization

### 2. **ConcurrentProgrammingInGo**
Concurrent programming patterns and techniques from the Pluralsight course.

**Topics Covered:**
- Goroutines and concurrent execution
- Channels for communication
- Synchronization primitives
- Order management system example (`order.go`)

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ installed
- Basic familiarity with command-line tools

### Running Examples
```bash
# Run a specific example from Chapter 1
cd ALinguagemDeProgramacaoGo
go run chapter1/echo1.go arg1 arg2

# Run concurrent programming examples
cd ConcurrentProgrammingInGo
go run main.go
```

## 📝 Notes
- All code examples are educational and designed for learning purposes
- Each subdirectory contains its own `go.mod` file for dependency management
- Examples follow Go conventions and best practices

## 🔗 Additional Resources
- [Official Go Documentation](https://golang.org/doc)
- [Go by Example](https://gobyexample.com)
- [Effective Go](https://golang.org/doc/effective_go)