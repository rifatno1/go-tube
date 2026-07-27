package layouts

import (
	"go-tube/internal/app/model"
	"go-tube/internal/app/view/components"
)

func Error(m model.Model, strArr *[]string) {
	*strArr = append(*strArr, components.Header(m))

	switch m.ErrorType {
	case model.ErrorType_BinError:
		width := 80
		contents := Error_BinError(m, width)
		*strArr = append(*strArr, components.BoxCenter(m, width, contents))

	// model.ErrorType_RootDirError
	// model.ErrorType_GetSettingsError
	// model.ErrorType_None
	default:
		width := 80
		contents := Error_General(m, width)
		*strArr = append(*strArr, components.BoxCenter(m, width, contents))
	}
}
