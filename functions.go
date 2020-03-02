package templateutils

import (
	"strings"
	"text/template"
)

var funcMap = map[string]interface{}{
	// String
	"toCamel":   ToCamel,
	"toKebab":   ToKebab,
	"toLower":   strings.ToLower,
	"toPascal":  ToPascal,
	"toSnake":   ToSnack,
	"toUpper":   strings.ToUpper,
	"zeroValue": ZeroValue,

	// Utils
	"in": In,
}

// FuncMap will return the template utils as FuncMap.
func FuncMap() template.FuncMap {
	return funcMap
}
