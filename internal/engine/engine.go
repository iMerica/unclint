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

	suppressions := parseSuppressions(text)

	var findings []Finding

	for _, s := range sentences {
		for _, m := range e.matchers {
			f := m.Match(s)
			findings = append(findings, f...)
		}
	}

	for i := range findings {
		line, col := byteToLineCol(text, findings[i].StartByte)
		findings[i].File = path
		findings[i].Line = line
		findings[i].Column = col
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

		if isSuppressed(f, suppressions) {
			continue
		}

		filtered = append(filtered, f)
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

type Suppression struct {
	Type    string
	Line    int
	Targets []string
}

func parseSuppressions(text string) []Suppression {
	var suppressions []Suppression
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lineNum := i + 1
		idx := strings.Index(line, "unc-disable")
		if idx == -1 {
			continue
		}

		rem := line[idx:]
		parts := strings.Fields(rem)
		if len(parts) == 0 {
			continue
		}

		directive := parts[0]
		var targets []string
		for _, p := range parts[1:] {
			if strings.HasPrefix(p, "--") {
				break
			}
			targets = append(targets, p)
		}

		if directive == "unc-disable" {
			suppressions = append(suppressions, Suppression{Type: "file", Line: lineNum, Targets: targets})
		} else if directive == "unc-disable-next-line" {
			suppressions = append(suppressions, Suppression{Type: "next-line", Line: lineNum, Targets: targets})
		} else if directive == "unc-disable-line" {
			suppressions = append(suppressions, Suppression{Type: "line", Line: lineNum, Targets: targets})
		}
	}
	return suppressions
}

func byteToLineCol(text string, startByte int) (int, int) {
	line := 1
	col := 1
	for i := 0; i < startByte && i < len(text); i++ {
		if text[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return line, col
}

func isSuppressed(f Finding, suppressions []Suppression) bool {
	for _, s := range suppressions {
		match := len(s.Targets) == 0
		if !match {
			for _, t := range s.Targets {
				if t == f.RuleID || t == f.Category {
					match = true
					break
				}
			}
		}

		if !match {
			continue
		}

		if s.Type == "file" {
			return true
		}
		if s.Type == "line" && s.Line == f.Line {
			return true
		}
		if s.Type == "next-line" && s.Line+1 == f.Line {
			return true
		}
	}
	return false
}
