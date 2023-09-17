package pkg

import (
	"log"

	"gopkg.in/yaml.v3"
)

// Dependencies reflects the schema for the tool-versions file
type Dependencies struct {
	Data map[string]map[string]string `yaml:",inline"`
}

// ReadTools reads from a byteslice of a tool-versions compatible file into d
func ReadTools(toolsSpec []byte, d *Dependencies) {
	if err := yaml.Unmarshal(toolsSpec, &d); err != nil {
		log.Fatal(err)
	}
}
