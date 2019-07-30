package gtranslate

import (
	"golang.org/x/text/language"
	"time"
)

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From  string
	To    string
	Tries int
	Delay time.Duration
}

// Translate translate a text using native tags offer by go language
func Translate(text string, from language.Tag, to language.Tag) (string, error) {
	translated, err := translate(text, from.String(), to.String(), false, 2, 0)
	if err != nil {
		return "", err
	}

	return translated, nil
}

// TranslateWithParams translate a text with simple params as string
func TranslateWithParams(text string, params TranslationParams) (string, error) {
	translated, err := translate(text, params.From, params.To, true, params.Tries, params.Delay)
	if err != nil {
		return "", err
	}

	return translated, nil
}
