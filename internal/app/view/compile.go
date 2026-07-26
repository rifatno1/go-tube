package view

import "github.com/charmbracelet/lipgloss"

func Compile(strArr []string, width, height int) string {
	return lipgloss.Place(
		width,
		height,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Left, strArr...),
	)
}
