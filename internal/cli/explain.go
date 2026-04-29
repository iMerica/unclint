package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yourname/unc/internal/config"
	"github.com/yourname/unc/internal/engine"
	"github.com/yourname/unc/internal/output"
	"github.com/yourname/unc/internal/rules"
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
