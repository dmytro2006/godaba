# Godaba — Golang Database Viewer (TUI)

**Godaba** (GOlang DAtaBAse viewer) is a terminal user interface app for exploring SQL databases, running queries, and exporting results to CSV — built entirely in Go.

Designed as a keyboard-first developer tool, Godaba provides a clean and fast way to inspect databases directly from the terminal.

## Features

- Browse databases and tables
- View table data with pagination
- Execute custom SQL queries
- Pretty tabular result view
- Export query results to CSV
- Keyboard-driven workflow
- Clean split-pane terminal UI

## Supported DBs

- [ ] SQLite
- [ ] PostgreSQL
- [ ] MySQL
  
## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework fo Go
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components

## Getting Started

```bash
git clone https://github.com/yourusername/godaba
cd godaba
go run cmd/main.go
```

## Architecture

```
/cmd
/internal
  /ui       → Bubble Tea model, update, view
  /db       → Database connection and queries
  /styles   → Lip Gloss styling
```

## License

MIT