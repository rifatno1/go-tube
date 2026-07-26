package init

import (
	"go-tube/internal/app/handlers"
	"go-tube/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

func Init(model.Model) tea.Cmd {
	return tea.Batch(
		handlers.Trigger_lib_GetDir(),
	)
}
