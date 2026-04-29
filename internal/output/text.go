package output

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/iMerica/unclint/internal/engine"
)

var (
	fileStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	ruleStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	severityStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	textStyle       = lipgloss.NewStyle().Italic(true)
	suggestionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	scoreStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("220"))
	passStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("46"))
	failStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196"))
)

func PrintText(w io.Writer, results []engine.Result, maxScore int, noColor bool) {
	if noColor {
		// A simple disable for MVP
		lipgloss.SetColorProfile(0)
	}

	for _, r := range results {
		if len(r.Findings) == 0 {
			continue
		}

		for _, f := range r.Findings {
			sev := "warning"
			if f.Severity == 2 {
				sev = "error"
			}

			fmt.Fprintf(w, "%s:%d:%d  %s  %s  weight %d\n",
				fileStyle.Render(f.File),
				f.Line, f.Column,
				ruleStyle.Render(f.RuleID),
				severityStyle.Render(sev),
				f.Weight,
			)

			fmt.Fprintf(w, "\"%s\"\n\n", textStyle.Render(f.Text))
			fmt.Fprintf(w, "Problem:\n%s\n\n", f.Message)

			if f.Suggestion != "" {
				fmt.Fprintf(w, "Suggestion:\n%s\n\n", suggestionStyle.Render(f.Suggestion))
			}
		}

		fmt.Fprintf(w, "Score: %s\n", scoreStyle.Render(fmt.Sprint(r.Score)))
		if r.Score > maxScore {
			fmt.Fprintf(w, "Result: %s\n\n", failStyle.Render("FAIL"))
		} else {
			fmt.Fprintf(w, "Result: %s\n\n", passStyle.Render("PASS"))
		}
	}
}
