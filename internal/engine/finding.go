package engine

type Finding struct {
	File       string   `json:"file,omitempty"`
	Line       int      `json:"line"`
	Column     int      `json:"column"`
	EndLine    int      `json:"end_line,omitempty"`
	EndColumn  int      `json:"end_column,omitempty"`
	RuleID     string   `json:"rule_id"`
	Category   string   `json:"category"`
	Severity   int      `json:"severity"`
	Weight     int      `json:"weight"`
	Message    string   `json:"message"`
	Suggestion string   `json:"suggestion,omitempty"`
	Text       string   `json:"text"`
	Tags       []string `json:"tags,omitempty"`
}
