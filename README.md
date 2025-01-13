# Go Todo CLI

This project is a command-line interface (CLI) application for managing a todo list, built with Go.

## Purpose

The purpose of this project is to provide a simple and efficient way to manage your tasks from the command line. It is designed to be a practical tool for developers and users who prefer using the terminal for their daily tasks.

## Features

- Add new tasks
- List all tasks
- Mark tasks as completed or in progress
- Delete tasks
- Filter tasks by status

## Installation

To install the Go Todo CLI, follow these steps:

1. Clone the repository:
  ```sh
  git clone https://github.com/H4zlq/go-todo-cli.git
  ```
2. Navigate to the project directory:
  ```sh
  cd go-todo-cli
  ```
3. Build the application:
  ```sh
  go build -o todo
  ```
4. Move the executable to a directory in your PATH:
  ```sh
  mv todo /usr/local/bin/
  ```

## Usage

Here are some basic commands to get you started:

- Add a new task:
  ```sh
  todo add "Buy groceries"
  ```
- List all tasks:
  ```sh
  todo list
  ```
- Mark a task as completed:
  ```sh
  todo mark-done 1
  ```
- Mark a task as in progress:
  ```sh
  todo mark-in-progress 1
  ```
- Delete a task:
  ```sh
  todo delete 1
  ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## Roadmap

- [ ] Add support for due dates
- [ ] Implement task prioritization
- [ ] Add search functionality
- [ ] Create a web interface
- [ ] Improve test coverage

## Contact

For any questions or suggestions, please open an issue or contact the project maintainer at [hazinaim72@gmail.com](mailto:hazinaim72@gmail.com).
