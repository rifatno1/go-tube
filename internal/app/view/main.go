package view

import (
	"go-tube/internal/app/model"
	"go-tube/internal/app/view/layouts"
)

func View(m model.Model) string {
	strArr := []string{}

	switch m.Active_layout {
	case model.Layout_Main:
		layouts.Main(m, &strArr)
	case model.Layout_Settings:
		layouts.Settings(m, &strArr)
	case model.Layout_Error:
		layouts.Error(m, &strArr)
	default:
		strArr = append(strArr, "Unknown layout")
	}

	return Compile(strArr, m.Width, m.Height)
}
