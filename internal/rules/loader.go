package rules

import (
	"fmt"
	"path/filepath"

	rulespack "github.com/iMerica/unclint/rules"
	"gopkg.in/yaml.v3"
)

type RulesFile struct {
	Rules []Rule `yaml:"rules"`
}

func LoadDefaultRules() ([]Rule, error) {
	var allRules []Rule

	entries, err := rulespack.FS.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded rules: %w", err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".yml" {
			continue
		}

		data, err := rulespack.FS.ReadFile(entry.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read rule file %s: %w", entry.Name(), err)
		}

		var rf RulesFile
		if err := yaml.Unmarshal(data, &rf); err != nil {
			return nil, fmt.Errorf("failed to parse rule file %s: %w", entry.Name(), err)
		}

		allRules = append(allRules, rf.Rules...)
	}

	return allRules, nil
}
