package ui

import (
	"godaba/internal/styles"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	// builder := strings.Builder{}
	// builder.WriteString("Godaba\n")
	// builder.WriteString(m.help.View(m.keys))
	sidebarStyle := styles.AreaStyle
	switch m.mode.(type) {
	case TableSelectingMode:
		sidebarStyle = sidebarStyle.BorderForeground(lipgloss.Red)
	}

	sidebar := sidebarStyle.Height(m.height - 1).Width(30).Render(m.sidebar.View())

	view := tea.NewView(sidebar + "\n" + m.help.View(m.keys))
	view.AltScreen = true
	return view
}
