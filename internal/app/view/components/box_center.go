package components

import (
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func BoxCenter(m model.Model, width int, contents func(m model.Model, width int) []string) string {
	width = min(m.Width-2, width)
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(width)
	rendered := style.Render(
		lipgloss.JoinVertical(lipgloss.Center, contents(m, width)...),
	)

	return lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, rendered)
}
