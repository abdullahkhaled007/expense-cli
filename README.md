# expensr (expense-cli)

A small, beginner-friendly command-line expense tracker written in Go. Use this project to learn Go, practice writing services, stores and CLI handlers, and run a simple persistent expense store backed by a JSON file.

This repository is intentionally small and focused on learning: it demonstrates a domain model (`Expense`), an in-memory + file-backed `Store`, a `Service` layer, unit tests, and a thin `cmd` CLI layer.

## Quickstart

Open a PowerShell terminal and run from the project root (where `go.mod` is):

```powershell
cd 'C:\Users\a.khaled\IdeaProjects\golang'

# Run a one-off command (uses a JSON file `expenses.json` in the working dir)
go run . add "Coffee" 3.25 "morning"
go run . list
go run . delete 1

# Or build and run the executable
go build -o expensr .
.\expensr.exe add "Lunch" 12.5 "food"
.\expensr.exe list
```

Notes:
- The CLI stores data in `expenses.json` by default (in the current working directory). The file is meant to be local runtime data and should typically be ignored in your git repo.
- If you want to keep sample/demo data, create `expenses.example.json` and commit that instead.

## Commands (examples)

- `add <name> <price> [description]` — add a new expense. Example: `add "Coffee" 3.25 "morning"`
- `list` — list saved expenses (ID, name, price, date)
- `delete <id>` — delete expense by numeric ID

All commands are implemented as small, testable handlers that accept `[]string` and a `*expense.Service` so you can call them from tests easily.

## Tests and Coverage

Run tests for the repo:

```powershell
go test ./...
```

Get a coverage profile and view a human-readable summary:

```powershell
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Create an HTML coverage report and open it:

```powershell
go tool cover -html=coverage.out -o coverage.html
Start-Process .\coverage.html
```

## Development notes

- The service constructor `NewServiceWithFile(path string)` creates a `Service` backed by a JSON file. The default used by the CLI is `expenses.json`.
- Persistence is implemented by writing the full list to disk after add/delete operations. This is simple and easy to inspect but not optimized for large datasets. For learning purposes, this is sufficient.
- The CLI is intentionally thin: it parses arguments and calls service methods
