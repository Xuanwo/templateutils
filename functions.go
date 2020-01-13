package templateutils

import (
	"strings"
	"text/template"
)

var funcMap = map[string]interface{}{
	"toCamel":   ToCamel,
	"toKebab":   ToKebab,
	"toLower":   strings.ToLower,
	"toPascal":  ToPascal,
	"toSnake":   ToSnack,
	"toUpper":   strings.ToUpper,
	"zeroValue": ZeroValue,
}

// FuncMap will return the template utils as FuncMap.
func FuncMap() template.FuncMap {
	return funcMap
}
