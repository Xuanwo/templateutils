package templateutils

import (
	"testing"
	"text/template"
)

func TestFuncMap(t *testing.T) {
	v := FuncMap()
	if v == nil {
		t.Errorf("get func map failed.")
	}

	x := template.New("test").Funcs(v)
	if x == nil {
		t.Errorf("apply func map to text/template.")
	}
}
