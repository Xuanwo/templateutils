package templateutils

import (
	"text/template"

	"github.com/huandu/xstrings"
)

var funcMap = map[string]interface{}{
	"camelCase": xstrings.ToCamelCase,
	"zeroValue": ZeroValue,
}

// FuncMap will return the template utils as FuncMap.
func FuncMap() template.FuncMap {
	return funcMap
}
