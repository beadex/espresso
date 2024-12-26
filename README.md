# Espresso

This project is a terminal-based application built with Go that provides a simple GUI and backend functionality. It uses SQLite3 for data storage and management.

## Project Structure

```
espresso
├── cmd
│   └── main.go          # Entry point of the application
├── lib
│   ├── backend
│   │   └── backend.go   # Backend logic and business handling
│   ├── gui
│   │   └── gui.go       # Terminal-based GUI management
│   └── database
│       └── database.go   # Database interactions and queries
├── go.mod                # Module definition and dependencies
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**

   ```
   git clone <repository-url>
   cd espresso
   ```

2. **Install dependencies:**

   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage Guidelines

- The application starts with a terminal-based GUI that allows users to interact with the backend.
- Follow the prompts in the terminal to navigate through the application.
- Ensure that SQLite3 is installed on your system for database functionality.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
