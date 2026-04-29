package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Version     int            `mapstructure:"version"`
	Context     string         `mapstructure:"context"`
	MaxScore    int            `mapstructure:"max_score"`
	MinSeverity int            `mapstructure:"min_severity"`
	Include     []string       `mapstructure:"include"`
	Exclude     []string       `mapstructure:"exclude"`
	Rules       map[string]any `mapstructure:"rules"`
	Overrides   []Override     `mapstructure:"overrides"`
	Allow       Allow          `mapstructure:"allow"`
	Disable     []string       `mapstructure:"disable"`
}

type Override struct {
	Path     string `mapstructure:"path"`
	Context  string `mapstructure:"context"`
	MaxScore int    `mapstructure:"max_score"`
}

type Allow struct {
	Terms []string `mapstructure:"terms"`
}

func Load(cfgFile string) (*Config, error) {
	v := viper.New()
	SetDefaults(v)

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		// Default search paths
		cwd, err := os.Getwd()
		if err == nil {
			v.AddConfigPath(cwd)
		}
		v.SetConfigName(".uncrc")
		v.SetConfigType("yml")
	}

	// It's ok if the config file isn't found when using defaults
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok && cfgFile != "" {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	return &cfg, nil
}
