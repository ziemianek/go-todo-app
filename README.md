# Go Todo App

## Description

This is a simple command-line todo application written in Go. It's designed as a learning project to explore various aspects of the Go programming language, with a focus on file handling for data persistence.

## Features

- Add new todo items
- List all todos
- Mark todos as completed
- Delete todos
- Save todos to a file
- Load todos from a file
- Filter todos (show completed/incomplete)
- Basic CLI interface
- Search functionality

## Project Structure
```
go-todo-app/
│
├── main.go
├── todo/
│ ├── todo.go
│ └── file_handler.go
├── cli/
│ └── cli.go
├── data/
│ └── todos.json
```

## Features

- Add new todo items
- List all todos
- Mark todos as completed
- Delete todos
- Save todos to a file
- Load todos from a file
- Filter todos (show completed/incomplete)
- Basic CLI interface
- Search functionality

## Getting Started

### Prerequisites

- Go 1.20.4
- Make

### Installation

1. Clone the repository:
```sh
git clone https://github.com/yourusername/go-todo-app.git

```

2. Navigate to the project directory:
```sh
cd go-todo-app
```

3. Build the app:

```sh
make build
```

### Usage

Run the application:

```sh
./go-todo-app

```

**Follow the on-screen prompts to interact with the todo list.**

## Learning Objectives

This project covers the following Go concepts:

- Structs and methods
- Slices and maps
- File I/O operations
- Error handling
- Date and time manipulation
- Command-line argument parsing
- Basic string manipulation and searching

## Contributing

This is a personal learning project, and contributions are not expected. However, suggestions for improvements or bug reports are welcome.

## License

This project is open source and available under the [MIT License](LICENSE).
