package model

type layout struct {
	value string
}

func (l layout) String() string {
	return l.value
}

var (
	Layout_Error    = layout{value: "error"}
	Layout_Settings = layout{value: "settings"}
	Layout_Main     = layout{value: "main"}
)
