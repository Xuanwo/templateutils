package templateutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroValue(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"int", "0"},
		{"int64", "0"},
		{"float64", "0"},
		{"complex64", "0"},
		{"string", `""`},
		{"bool", "false"},
		{"map[string]int64", "map[string]int64{}"},
		{"struct { Name string }", "struct { Name string }{}"},
		{"os.File", "os.File{}"},
		{"[]int64", "nil"},
		{"*int64", "nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroValue(tt.name); got != tt.want {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitStringViaSpecialChars(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"abc", []string{"abc"}},
		{"a-bc", []string{"a", "bc"}},
		{"a,bc", []string{"a", "bc"}},
		{"ABC", []string{"ABC"}},
		{"A-BC", []string{"A", "BC"}},
		{"a--b", []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitStringViaSpecialChars(tt.name)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestSplitStringViaUpperChars(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"abc", []string{"abc"}},
		{"Abc", []string{"Abc"}},
		{"AbC", []string{"Ab", "C"}},
		{"AbCd", []string{"Ab", "Cd"}},
		{"AAA", []string{"AAA"}},
		{"AAABb", []string{"AAA", "Bb"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitStringViaUpperChars(tt.name)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
