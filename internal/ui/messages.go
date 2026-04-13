package ui

import tea "charm.land/bubbletea/v2"

type TablesLoadedMsg struct {
	tables []string
}

type TableSelectedMsg struct{}

type ErrorMsg struct{}

func TablesLoaded(tables []string) tea.Msg {
	return TablesLoadedMsg{tables: tables}
}

func TableSelected() tea.Msg {
	return TableSelectedMsg{}
}
