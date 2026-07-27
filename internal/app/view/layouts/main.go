package layouts

import (
	"go-tube/internal/app/model"
	"go-tube/internal/app/view/components"
)

func Main(m model.Model, strArr *[]string) {
	*strArr = append(*strArr, components.Header(m))
	*strArr = append(
		*strArr,
		components.BoxCenter(m, m.Width, Main_Search),
	)
}
