package rules

type Rule struct {
	ID         string   `yaml:"id"`
	Kind       string   `yaml:"kind"`
	Category   string   `yaml:"category"`
	Pattern    string   `yaml:"pattern,omitempty"`
	Terms      []string `yaml:"terms,omitempty"`
	Lemma      string   `yaml:"lemma,omitempty"`
	POS        string   `yaml:"pos,omitempty"`
	Near       []string `yaml:"near,omitempty"`
	Window     int      `yaml:"window,omitempty"`
	Severity   int      `yaml:"severity"`
	Weight     int      `yaml:"weight"`
	Message    string   `yaml:"message"`
	Suggestion string   `yaml:"suggestion,omitempty"`
	Contexts   []string `yaml:"contexts,omitempty"`
}
