package rules

import (
	"strings"

	"github.com/iMerica/unclint/internal/engine"
	"github.com/iMerica/unclint/internal/nlp"
)

type Matcher interface {
	Match(sentence nlp.AnalyzedSentence) []engine.Finding
}

func BuildMatchers(rules []Rule) []engine.RuleMatcher {
	var matchers []engine.RuleMatcher
	for _, r := range rules {
		switch r.Kind {
		case "phrase":
			matchers = append(matchers, NewPhraseMatcher(r))
		case "lemma_pos":
			matchers = append(matchers, NewLemmaPOSMatcher(r))
		case "construction":
			matchers = append(matchers, NewConstructionMatcher(r))
		}
	}
	return matchers
}

func makeFinding(r Rule, sentence nlp.AnalyzedSentence) engine.Finding {
	return engine.Finding{
		StartByte:  sentence.StartByte,
		RuleID:     r.ID,
		Category:   r.Category,
		Severity:   r.Severity,
		Weight:     r.Weight,
		Message:    r.Message,
		Suggestion: r.Suggestion,
		Text:       sentence.Text,
	}
}

type PhraseMatcher struct {
	rule    Rule
	pattern string
}

func NewPhraseMatcher(r Rule) *PhraseMatcher {
	return &PhraseMatcher{
		rule:    r,
		pattern: strings.ToLower(r.Pattern),
	}
}

func (m *PhraseMatcher) Match(sentence nlp.AnalyzedSentence) []engine.Finding {
	lower := strings.ToLower(sentence.Text)
	if strings.Contains(lower, m.pattern) {
		return []engine.Finding{makeFinding(m.rule, sentence)}
	}
	return nil
}
