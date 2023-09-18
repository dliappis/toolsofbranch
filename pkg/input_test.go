package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTools(t *testing.T) {
	inputTest1WithEmptyCommon := `
versioned:
  toolFamily1:
    tool1: ">=4.2.3"
    tool2: "<4.2.3"
`

	var parsedYaml Dependencies
	wantsTest1 := Dependencies{
		Versioned: map[string]map[string]string{
			"toolFamily1": {
				"tool1": ">=4.2.3",
				"tool2": "<4.2.3",
			},
		},
	}
	ReadTools([]byte(inputTest1WithEmptyCommon), &parsedYaml)
	assert.Equal(t, parsedYaml, wantsTest1, "Test 1 failed to parse yaml file into expected content")

	inputTest2WithoutCommon := `
versioned:
  toolFamily1:
    tool1: ">10.11.3"
    tool2: ">7.9, <10.3.1"
`

	wantsTest2 := Dependencies{
		Versioned: map[string]map[string]string{
			"toolFamily1": {
				"tool1": ">10.11.3",
				"tool2": ">7.9, <10.3.1",
			},
		},
	}
	ReadTools([]byte(inputTest2WithoutCommon), &parsedYaml)
	assert.Equal(t, parsedYaml, wantsTest2, "Test 2 failed to parse yaml file into expected content")

	inputTest3WithCommon := `
common:
  - yq
  - jq
versioned:
  toolFamily1:
    tool1: ">10.11.3"
    tool2: ">7.9, <10.3.1"
`

	wantsTest3 := Dependencies{
		Common: []string{"yq", "jq"},
		Versioned: map[string]map[string]string{
			"toolFamily1": {
				"tool1": ">10.11.3",
				"tool2": ">7.9, <10.3.1",
			},
		},
	}
	ReadTools([]byte(inputTest3WithCommon), &parsedYaml)
	assert.Equal(t, parsedYaml, wantsTest3, "Test 3 failed to parse yaml file into expected content")
}
