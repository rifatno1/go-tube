package layouts

import (
	"go-tube/internal/app/model"
	"go-tube/internal/app/view/components"
)

func Error(m model.Model, strArr *[]string) {
	*strArr = append(*strArr, components.Header(m))

	switch m.ErrorType {
	case model.ErrorType_BinError:
		*strArr = append(
			*strArr,
			components.BoxCenter(m, 80, Error_BinError),
		)

	// model.ErrorType_RootDirError
	// model.ErrorType_GetSettingsError
	// model.ErrorType_None
	default:
		*strArr = append(
			*strArr,
			components.BoxCenter(m, 80, Error_General),
		)
	}
}
