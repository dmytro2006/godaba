package main

import (
	"godaba/internal/ui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	m := ui.NewModel()
	program := tea.NewProgram(m)
	program.Run()
}
