package templateutils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource_ParseContent(t *testing.T) {
	t.Run("parse interface", func(t *testing.T) {
		filename := "testdata/source_1"
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("read file: %v", err)
		}

		s, err := ParseContent(filename, content)
		if err != nil {
			t.Errorf("parse content: %v", err)
		}

		assert.Equal(t, 7, len(s.Interfaces))
	})

	t.Run("parse functions", func(t *testing.T) {
		filename := "testdata/source_2"
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("read file: %v", err)
		}

		s, err := ParseContent(filename, content)
		if err != nil {
			t.Errorf("parse content: %v", err)
		}

		for _, v := range s.Methods {
			println(v.Name)
		}
	})
}
