package templateutils

import (
	"os"
	"testing"
)

func TestZeroValue(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want string
	}{
		{"integer", 123, "0"},
		{"float", 3.15, "0"},
		{"string", "Hello, World!", ""},
		{"bool", true, "false"},
		{"map", map[string]int64{"a": 1}, "map[string]int64{}"},
		{"struct", struct{ Name string }{"xxxx"}, "struct { Name string }{}"},
		{"name struct", os.File{}, "os.File{}"},
		{"slice", []int64{1, 2, 3}, "nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroValue(tt.args); got != tt.want {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
