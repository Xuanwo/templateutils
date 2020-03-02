package templateutils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	tests := []struct {
		name       string
		inputItem  interface{}
		inputValue interface{}
		want       bool
		wantErr    error
	}{
		{
			"string in array",
			[1]string{"x"},
			"x",
			true,
			nil,
		},
		{
			"int in array",
			[1]int{1},
			1,
			true,
			nil,
		},
		{
			"string not in array",
			[1]string{"x"},
			"y",
			false,
			nil,
		},
		{
			"int not in array",
			[1]int{1},
			2,
			false,
			nil,
		},
		{
			"string in string",
			"abc",
			"a",
			true,
			nil,
		},
		{
			"string not in string",
			"abc",
			"d",
			false,
			nil,
		},
		{
			"string in string map",
			map[string]string{
				"a": "x",
			},
			"a",
			true,
			nil,
		},
		{
			"string not in string map",
			map[string]string{
				"a": "x",
			},
			"b",
			false,
			nil,
		},
		{
			"string in bool map",
			map[string]bool{
				"a": false,
			},
			"b",
			false,
			nil,
		},
		{
			"string in int map",
			map[int]bool{
				1: false,
			},
			"b",
			false,
			errors.New("test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := In(reflect.ValueOf(tt.inputItem), reflect.ValueOf(tt.inputValue))
			if tt.wantErr != nil {
				assert.NotNil(t, gotErr)
				return
			}

			assert.Nil(t, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
