package output

import (
	"fmt"
	"io"

	"github.com/yourname/unc/internal/engine"
)

func PrintGitHub(w io.Writer, results []engine.Result) {
	for _, r := range results {
		for _, f := range r.Findings {
			level := "warning"
			if f.Severity == 2 {
				level = "error"
			}
			fmt.Fprintf(w, "::%s file=%s,line=%d,col=%d::%s: %s\n",
				level,
				f.File,
				f.Line,
				f.Column,
				f.RuleID,
				f.Message,
			)
		}
	}
}
