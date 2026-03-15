package scenario

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFile(path string) (*Scenario, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s Scenario
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}

	if s.Name == "" {
		return nil, fmt.Errorf("scenario name is required")
	}

	return &s, nil
}
