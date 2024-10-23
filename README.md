# Command Executor

This project is a simple command execution application that exposes an HTTP API to perform network ping commands and retrieve system information.

## How to Build

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/command-executor.git
    cd command-executor
    ```

2. **Initialize the Go module (if necessary):**

    ```bash
    go mod init github.com/yourusername/command-executor
    go mod tidy
    ```

3. **Build the application:**

    ```bash
    go build -o command-executor main.go
    ```

    This will create a `command-executor` binary that you can run.

## How to Run

To run the application and expose the API on port `8081`:

```bash
./command-executor

curl -X POST http://localhost:8081/exec -d '{"type":"ping", "data":"google.com"}' -H "Content-Type: application/json"

curl -X POST http://localhost:8081/exec -d '{"type":"sysinfo"}' -H "Content-Type: application/json"

