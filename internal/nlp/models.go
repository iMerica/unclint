package nlp

type AnalyzedSentence struct {
	Text      string
	StartByte int
	EndByte   int
	Tokens    []Token
	Features  Features
}

type Token struct {
	Text      string
	Lower     string
	Lemma     string
	POS       string
	Entity    string
	StartByte int
	EndByte   int
}

type Features struct {
	TokenCount              int
	WordCount               int
	SlangCount              int
	CorporateJargonCount    int
	AbstractNounCount       int
	NominalizationCount     int
	IntensifierCount        int
	HedgeCount              int
	ExclamationCount        int
	QuestionCount           int
	SlangDensity            float64
	CorporateDensity        float64
	AbstractNounDensity     float64
	HasMetaSlang            bool
	HasBoomerFraming        bool
	HasCorporateNounPile    bool
	HasJargonVerbObjectPair bool
}
