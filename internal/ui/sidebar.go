package ui

import (
	"fmt"
	"io"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type item struct {
	Value any
}

func (i item) FilterValue() string {
	return fmt.Sprint(i.Value)
}

// func (i item) Title() string {
// 	return fmt.Sprint(i.value)
// }

// func (i item) Description() string {
// 	return ""
// }

type sidebarDelegate struct {
}

func newSidebarDelegate() sidebarDelegate {
	return sidebarDelegate{}
}

func (d sidebarDelegate) Height() int {
	return 1
}

func (d sidebarDelegate) Spacing() int {
	return 0
}

func (d sidebarDelegate) Render(w io.Writer, m list.Model, index int, it list.Item) {
	i, ok := it.(item)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprintf(w, ">%s", i.Value)
		return
	}
	fmt.Fprintf(w, " %s", i.Value)
}

func (d sidebarDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}
