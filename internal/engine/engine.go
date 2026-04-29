package engine

import (
	"strings"

	"github.com/yourname/unc/internal/config"
	"github.com/yourname/unc/internal/nlp"
)

type RuleMatcher interface {
	Match(sentence nlp.AnalyzedSentence) []Finding
}

type Engine struct {
	config   *config.Config
	matchers []RuleMatcher
	triggers map[string]struct{}
}

func New(cfg *config.Config, matchers []RuleMatcher) *Engine {
	// A simple set of triggers for fast mode
	triggers := map[string]struct{}{
		"leverage": {}, "utilize": {}, "operationalize": {}, "socialize": {},
		"drive": {}, "unlock": {}, "circle back": {}, "touch base": {},
		"move the needle": {}, "low-hanging fruit": {},
	}

	return &Engine{
		config:   cfg,
		matchers: matchers,
		triggers: triggers,
	}
}

func (e *Engine) LintText(path string, text string, deep bool) (Result, error) {
	sentences, err := nlp.Analyze(text, deep, e.triggers)
	if err != nil {
		return Result{}, err
	}

	var findings []Finding

	for _, s := range sentences {
		for _, m := range e.matchers {
			f := m.Match(s)
			findings = append(findings, f...)
		}
	}

	// Filter findings based on config rules enabled/disabled
	var filtered []Finding
	for _, f := range findings {
		// skip disabled categories or specific rules
		enabled, ok := e.config.Rules[f.Category]
		if ok {
			if v, isBool := enabled.(bool); isBool && !v {
				continue // disabled by category
			}
		}

		disabledRule := false
		for _, d := range e.config.Disable {
			if d == f.RuleID || d == f.Category {
				disabledRule = true
				break
			}
		}
		if disabledRule {
			continue
		}

		// check allows terms
		allowed := false
		for _, term := range e.config.Allow.Terms {
			if strings.Contains(strings.ToLower(f.Text), strings.ToLower(term)) {
				allowed = true
				break
			}
		}
		if allowed {
			continue
		}

		filtered = append(filtered, f)
	}

	// Filter by inline suppressions here (Phase 8 logic can be added later)

	// Update finding file/line info
	for i := range filtered {
		filtered[i].File = path
		filtered[i].Line = 1 // basic for MVP, needs byte-to-line mapping
		filtered[i].Column = 1
	}

	score := Score(filtered)
	pass := score <= e.config.MaxScore

	return Result{
		Path:     path,
		Score:    score,
		Pass:     pass,
		Findings: filtered,
	}, nil
}
