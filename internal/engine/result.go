package engine

type Result struct {
	Path     string    `json:"path,omitempty"`
	Score    int       `json:"score"`
	Pass     bool      `json:"pass"`
	Findings []Finding `json:"findings"`
}
