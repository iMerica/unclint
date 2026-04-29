package rules

import (
	"strings"

	"github.com/yourname/unc/internal/engine"
	"github.com/yourname/unc/internal/nlp"
)

type LemmaPOSMatcher struct {
	rule  Rule
	lemma string
	pos   string
}

func NewLemmaPOSMatcher(r Rule) *LemmaPOSMatcher {
	return &LemmaPOSMatcher{
		rule:  r,
		lemma: strings.ToLower(r.Lemma),
		pos:   r.POS,
	}
}

func (m *LemmaPOSMatcher) Match(sentence nlp.AnalyzedSentence) []engine.Finding {
	for _, token := range sentence.Tokens {
		if token.Lemma == m.lemma && strings.HasPrefix(token.POS, m.pos) {
			return []engine.Finding{makeFinding(m.rule, sentence)}
		}
	}
	return nil
}
