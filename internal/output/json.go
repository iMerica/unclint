package output

import (
	"encoding/json"
	"io"

	"github.com/iMerica/unclint/internal/engine"
)

type JSONOutput struct {
	Score int             `json:"score"`
	Pass  bool            `json:"pass"`
	Files []engine.Result `json:"files"`
}

func PrintJSON(w io.Writer, results []engine.Result, maxScore int) error {
	totalScore := 0
	pass := true

	for _, r := range results {
		totalScore += r.Score
		if r.Score > maxScore {
			pass = false
		}
	}

	out := JSONOutput{
		Score: totalScore,
		Pass:  pass,
		Files: results,
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
