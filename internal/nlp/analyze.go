package nlp

import (
	"strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/abadojack/whatlanggo"
	"github.com/jdkato/prose/v2"
	"github.com/neurosnap/sentences/english"
)

var lemmatizer *golem.Lemmatizer

func init() {
	var err error
	lemmatizer, err = golem.New(en.New())
	if err != nil {
		panic(err)
	}
}

func Analyze(text string, deep bool, triggers map[string]struct{}) ([]AnalyzedSentence, error) {
	info := whatlanggo.Detect(text)
	if info.Confidence > 0.9 && info.Lang != whatlanggo.Eng {
		return nil, nil // Skip non-English
	}

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		return nil, err
	}

	sentences := tokenizer.Tokenize(text)
	var analyzed []AnalyzedSentence

	offset := 0
	for _, s := range sentences {
		sentenceText := strings.TrimSpace(s.Text)
		if sentenceText == "" {
			continue
		}

		idx := strings.Index(text[offset:], sentenceText)
		startByte := offset
		if idx != -1 {
			startByte = offset + idx
		}

		as := AnalyzedSentence{
			Text:      sentenceText,
			StartByte: startByte,
			EndByte:   startByte + len(sentenceText),
		}
		offset = startByte + len(sentenceText)

		suspicious := IsSuspicious(sentenceText, triggers)
		if deep || suspicious {
			doc, err := prose.NewDocument(sentenceText)
			if err == nil {
				for _, t := range doc.Tokens() {
					lower := strings.ToLower(t.Text)
					lemma := lemmatizer.Lemma(lower)

					as.Tokens = append(as.Tokens, Token{
						Text:  t.Text,
						Lower: lower,
						Lemma: lemma,
						POS:   t.Tag,
					})
				}

				for _, ent := range doc.Entities() {
					for i, tok := range as.Tokens {
						if strings.Contains(ent.Text, tok.Text) {
							as.Tokens[i].Entity = ent.Label
						}
					}
				}
			}
		}

		analyzed = append(analyzed, as)
	}

	return analyzed, nil
}
