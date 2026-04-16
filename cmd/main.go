package main

import (
	"godaba/internal/db"
	"godaba/internal/ui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// conection := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=prefer", "localhost", "5432", "postgres", "12345678")
	connection := "postgres://postgres:12345678@localhost:5432?sslmode=prefer"
	pg, err := db.NewExplorer("postgres", connection)
	if err != nil {
		panic(err)
	}
	m := ui.NewModel(pg, ui.DefaultKeyMap)
	program := tea.NewProgram(m)
	program.Run()
}
