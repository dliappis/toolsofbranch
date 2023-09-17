package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTools(t *testing.T) {
	validToolsFile1 := `
toolFamily1:
  tool1: ">10.11.3"
  tool2: ">7.9, <10.3.1"`

	var deps Dependencies
	wants := Dependencies{
		Data: map[string]map[string]string{
			"toolFamily1": {
				"tool1": ">10.11.3",
				"tool2": ">7.9, <10.3.1",
			},
		},
	}
	ReadTools([]byte(validToolsFile1), &deps)

	assert.Equal(t, deps, wants, "Failed to parse yaml file into expected content")
}
