package ui

import (
	"godaba/internal/db"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textarea"
)

type Model struct {
	explorer db.Explorer

	mode Mode

	tables        []string
	selectedTable string

	sidebar list.Model
	table   table.Model
	editor  textarea.Model

	keys keyMap
	help help.Model

	view    string // "tables" | "rows" | "editor"
	loading bool
	err     error

	height int
	width  int
}

func NewModel(exp db.Explorer, keys keyMap) Model {
	m := Model{
		explorer: exp,
		keys:     keys,
		help:     help.New(),
		sidebar:  list.New([]list.Item{}, newSidebarDelegate(), 0, 0),
		mode:     TableSelectingMode{},
	}
	m.sidebar.SetShowHelp(false)
	m.sidebar.SetShowTitle(false)
	m.sidebar.SetShowStatusBar(false)
	m.sidebar.SetFilteringEnabled(false)
	return m
}
