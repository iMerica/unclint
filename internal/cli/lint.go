package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/iMerica/unclint/internal/config"
	"github.com/iMerica/unclint/internal/engine"
	"github.com/iMerica/unclint/internal/files"
	"github.com/iMerica/unclint/internal/output"
	"github.com/iMerica/unclint/internal/rules"
)

var (
	lintContext  string
	lintConfig   string
	lintInclude  []string
	lintExclude  []string
	lintFormat   string
	lintJSON     bool
	lintMaxScore int
	lintMinSev   int
	lintNoColor  bool
	lintQuiet    bool
	lintStdin    bool
	lintDeep     bool
)

var lintCmd = &cobra.Command{
	Use:   "lint [paths...]",
	Short: "Lint files for corporate and stale language",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(lintConfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Config error:", err)
			os.Exit(2)
		}

		// Apply flag overrides
		if cmd.Flags().Changed("max-score") {
			cfg.MaxScore = lintMaxScore
		}
		if cmd.Flags().Changed("min-severity") {
			cfg.MinSeverity = lintMinSev
		}

		allRules, err := rules.LoadDefaultRules()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Rule load error:", err)
			os.Exit(2)
		}

		matchers := rules.BuildMatchers(allRules)
		linter := engine.New(cfg, matchers)

		var results []engine.Result

		if lintStdin {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Stdin error:", err)
				os.Exit(2)
			}
			res, _ := linter.LintText("stdin", string(data), lintDeep)
			results = append(results, res)
		} else {
			includes := cfg.Include
			if len(lintInclude) > 0 {
				includes = lintInclude
			}
			excludes := cfg.Exclude
			if len(lintExclude) > 0 {
				excludes = lintExclude
			}

			matchedFiles, err := files.Discover(args, includes, excludes)
			if err != nil {
				fmt.Fprintln(os.Stderr, "File discovery error:", err)
				os.Exit(2)
			}

			for _, path := range matchedFiles {
				content, err := files.Read(path)
				if err != nil {
					continue
				}
				res, err := linter.LintText(path, content.Content, lintDeep)
				if err == nil {
					results = append(results, res)
				}
			}
		}

		pass := true
		for _, r := range results {
			if r.Score > cfg.MaxScore {
				pass = false
			}
		}

		if lintJSON {
			lintFormat = "json"
		}

		switch lintFormat {
		case "json":
			output.PrintJSON(os.Stdout, results, cfg.MaxScore)
		case "github":
			output.PrintGitHub(os.Stdout, results)
		default:
			output.PrintText(os.Stdout, results, cfg.MaxScore, lintNoColor)
		}

		if !pass {
			os.Exit(1)
		}
	},
}

func init() {
	lintCmd.Flags().StringVar(&lintContext, "context", "general", "brand, dating, social, docs, enterprise, sms, general")
	lintCmd.Flags().StringVar(&lintConfig, "config", "", "path to config file")
	lintCmd.Flags().StringSliceVar(&lintInclude, "include", nil, "glob patterns to include")
	lintCmd.Flags().StringSliceVar(&lintExclude, "exclude", nil, "glob patterns to exclude")
	lintCmd.Flags().StringVar(&lintFormat, "format", "text", "text, json, github")
	lintCmd.Flags().BoolVar(&lintJSON, "json", false, "shorthand for --format json")
	lintCmd.Flags().IntVar(&lintMaxScore, "max-score", 0, "max allowed score before exit 1")
	lintCmd.Flags().IntVar(&lintMinSev, "min-severity", 1, "minimum severity to report")
	lintCmd.Flags().BoolVar(&lintNoColor, "no-color", false, "disable color")
	lintCmd.Flags().BoolVar(&lintQuiet, "quiet", false, "only print summary")
	lintCmd.Flags().BoolVar(&lintStdin, "stdin", false, "read from stdin")
	lintCmd.Flags().BoolVar(&lintDeep, "deep", false, "run all NLP rules, not only suspicious sentences")

	rootCmd.AddCommand(lintCmd)
}
