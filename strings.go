package templateutils

import (
	"fmt"
	"reflect"
	"regexp"
)

// ZeroValue will return input data's default zero value.
func ZeroValue(in interface{}) string {
	typ := reflect.TypeOf(in)
	switch typ.Kind() {
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice:
		return "nil"
	case reflect.Map, reflect.Struct:
		r := regexp.MustCompile(`<(.*) Value>`)
		v := r.FindStringSubmatch(fmt.Sprintf("%v", reflect.Zero(typ).String()))
		return v[1] + "{}"
	default:
		return fmt.Sprintf("%v", reflect.Zero(typ))
	}
}
