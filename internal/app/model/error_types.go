package model

type ErrorType struct {
	value string
}

func (e ErrorType) String() string {
	return e.value
}

var (
	ErrorType_RootDirError     = ErrorType{value: "root_dir_error"}
	ErrorType_BinError         = ErrorType{value: "bin_error"}
	ErrorType_None             = ErrorType{value: ""}
	ErrorType_GetSettingsError = ErrorType{value: "get_settings_error"}
)
