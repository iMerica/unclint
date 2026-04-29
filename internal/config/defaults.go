package config

import "github.com/spf13/viper"

func SetDefaults(v *viper.Viper) {
	v.SetDefault("version", 1)
	v.SetDefault("context", "general")
	v.SetDefault("max_score", 0)
	v.SetDefault("min_severity", 1)
	v.SetDefault("include", []string{
		"**/*.txt",
		"**/*.md",
		"**/*.mdx",
		"**/*.html",
		"**/*.tsx",
		"**/*.jsx",
		"**/*.ts",
		"**/*.js",
	})
	v.SetDefault("exclude", []string{
		"node_modules/**",
		".next/**",
		"dist/**",
		"build/**",
		"coverage/**",
		".git/**",
		"vendor/**",
		"*.min.js",
		"*.lock",
		"package-lock.json",
		"pnpm-lock.yaml",
		"yarn.lock",
	})
	v.SetDefault("rules", map[string]bool{
		"corporate":     true,
		"stale":         true,
		"millennial":    true,
		"boomer":        true,
		"tryhard":       true,
		"meta_slang":    true,
		"internet_dead": true,
	})
	v.SetDefault("overrides", []Override{})
	v.SetDefault("allow", Allow{Terms: []string{}})
	v.SetDefault("disable", []string{})
}
