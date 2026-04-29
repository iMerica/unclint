package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/iMerica/unclint/internal/rules"
)

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Manage and view rules",
}

var rulesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print loaded rules",
	Run: func(cmd *cobra.Command, args []string) {
		allRules, _ := rules.LoadDefaultRules()
		for _, r := range allRules {
			if rulesCategory == "" || r.Category == rulesCategory {
				fmt.Printf("%s [%s] (weight: %d): %s\n", r.ID, r.Category, r.Weight, r.Message)
			}
		}
	},
}

var rulesCategory string

func init() {
	rulesListCmd.Flags().StringVar(&rulesCategory, "category", "", "Filter by category (e.g., corporate)")
	rulesCmd.AddCommand(rulesListCmd)
	rootCmd.AddCommand(rulesCmd)
}
