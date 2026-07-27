package handlers

import "go-tube/internal/app/model"

func setError(m *model.Model, errorType model.ErrorType, errorMessage string) {
	m.ErrorMessage = errorMessage
	m.ErrorType = errorType
	m.Active_layout = model.Layout_Error
}
func clearError(m *model.Model, errorType model.ErrorType) {
	if errorType == m.ErrorType {
		m.ErrorMessage = ""
		m.ErrorType = model.ErrorType_None
		m.Active_layout = model.Layout_Main
	}
}
