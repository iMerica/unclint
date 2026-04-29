package rules

import (
	"strings"

	"github.com/yourname/unc/internal/engine"
	"github.com/yourname/unc/internal/nlp"
)

type ConstructionMatcher struct {
	rule Rule
}

func NewConstructionMatcher(r Rule) *ConstructionMatcher {
	return &ConstructionMatcher{
		rule: r,
	}
}

var abstractNouns = map[string]bool{
	"alignment": true, "enablement": true, "stakeholder": true, "synergy": true,
	"transformation": true, "velocity": true, "roadmap": true, "impact": true,
	"outcomes": true, "learnings": true, "bandwidth": true, "visibility": true,
	"ecosystem": true, "flywheel": true, "moat": true, "paradigm": true,
	"engagement": true, "productivity": true, "potential": true, "value": true,
}

func (m *ConstructionMatcher) Match(sentence nlp.AnalyzedSentence) []engine.Finding {
	switch m.rule.Pattern {
	case "drive_abstract_noun":
		return m.matchDriveAbstract(sentence)
	case "unlock_abstract_noun":
		return m.matchUnlockAbstract(sentence)
	case "corporate_noun_pile":
		return m.matchNounPile(sentence)
	}
	return nil
}

func (m *ConstructionMatcher) matchDriveAbstract(sentence nlp.AnalyzedSentence) []engine.Finding {
	hasDrive := false
	for _, t := range sentence.Tokens {
		if t.Lemma == "drive" && strings.HasPrefix(t.POS, "VB") {
			hasDrive = true
			break
		}
	}
	if !hasDrive {
		return nil
	}

	for _, t := range sentence.Tokens {
		if abstractNouns[t.Lemma] || abstractNouns[t.Lower] {
			return []engine.Finding{makeFinding(m.rule, sentence)}
		}
	}
	return nil
}

func (m *ConstructionMatcher) matchUnlockAbstract(sentence nlp.AnalyzedSentence) []engine.Finding {
	hasUnlock := false
	for _, t := range sentence.Tokens {
		if t.Lemma == "unlock" && strings.HasPrefix(t.POS, "VB") {
			hasUnlock = true
			break
		}
	}
	if !hasUnlock {
		return nil
	}

	for _, t := range sentence.Tokens {
		if abstractNouns[t.Lemma] || abstractNouns[t.Lower] {
			return []engine.Finding{makeFinding(m.rule, sentence)}
		}
	}
	return nil
}

func (m *ConstructionMatcher) matchNounPile(sentence nlp.AnalyzedSentence) []engine.Finding {
	// A basic noun pile detector: 3 or more consecutive NNs
	consecutiveNouns := 0
	for _, t := range sentence.Tokens {
		if strings.HasPrefix(t.POS, "NN") {
			consecutiveNouns++
			if consecutiveNouns >= 3 {
				// We found a noun pile, could check if last is abstract noun but let's just flag it
				return []engine.Finding{makeFinding(m.rule, sentence)}
			}
		} else {
			consecutiveNouns = 0
		}
	}
	return nil
}
