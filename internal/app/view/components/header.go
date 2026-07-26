package components

import (
	"fmt"
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func Header(m model.Model) string {
	style := lipgloss.NewStyle().Width(m.Width).Align(lipgloss.Center).Background(m.Theme.BG_main).Bold(true).MarginBottom(1)
	if m.ErrorType == model.ErrorType_RootDirError {
		return style.Render(fmt.Sprintf("Go-Tube [Error] %dX%d", m.Width, m.Height))
	}
	return style.Render(fmt.Sprintf("Go-Tube [%s] %dX%d", m.Dir, m.Width, m.Height))
}
