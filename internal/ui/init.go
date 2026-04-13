package ui

import tea "charm.land/bubbletea/v2"

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		tables, _ := m.explorer.ListTables()
		return TablesLoadedMsg{tables: tables}
	}
}
