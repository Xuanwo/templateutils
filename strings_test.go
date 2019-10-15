package templateutils

import (
	"testing"
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
