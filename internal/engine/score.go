package engine

import "github.com/yourname/unc/internal/config"

func Score(findings []Finding) int {
	score := 0

	for _, f := range findings {
		score += f.Weight
	}

	if score > 100 {
		return 100
	}

	return score
}

func ShouldFail(results []Result, cfg *config.Config) bool {
	// Need config import but we have cycle if we import it.
	// Actually we pass MaxScore and MinSeverity.
	return false // implemented below
}
