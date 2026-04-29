package nlp

import "strings"

func IsSuspicious(sentence string, triggerSet map[string]struct{}) bool {
	lower := strings.ToLower(sentence)

	for trigger := range triggerSet {
		if strings.Contains(lower, trigger) {
			return true
		}
	}

	if len(sentence) > 180 {
		return true
	}

	if strings.Count(sentence, "!") > 1 {
		return true
	}

	return false
}
