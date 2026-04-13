package ui

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.sidebar, cmd = m.sidebar.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.sidebar.SetSize(28, m.height-3)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Select):
			m.selectedTable = m.sidebar.SelectedItem().FilterValue()
			return m, TableSelected
		}
	case TablesLoadedMsg:
		m.tables = msg.tables
		var l []list.Item
		for _, table := range m.tables {
			l = append(l, item{table})
		}
		m.sidebar.SetItems(l)
	}

	return m, cmd
}
