package cli

import (
	"os"

	"github.com/iMerica/unclint/internal/config"
	"github.com/iMerica/unclint/internal/engine"
	"github.com/iMerica/unclint/internal/output"
	"github.com/iMerica/unclint/internal/rules"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain [text]",
	Short: "Lint one string and show feature breakdown",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		cfg, _ := config.Load("")
		allRules, _ := rules.LoadDefaultRules()
		matchers := rules.BuildMatchers(allRules)
		linter := engine.New(cfg, matchers)

		res, _ := linter.LintText("cli", text, true) // force deep
		output.PrintText(os.Stdout, []engine.Result{res}, cfg.MaxScore, false)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
