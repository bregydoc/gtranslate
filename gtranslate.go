package gtranslate

import (
	"net/http"
	"time"

	"golang.org/x/text/language"
)

var GoogleHost = "google.com"

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
	Client     *http.Client
}

// Translate translate a text using native tags offer by go language
func Translate(text string, from language.Tag, to language.Tag, googleHost ...string) (string, error) {
	if len(googleHost) != 0 && googleHost[0] != "" {
		GoogleHost = googleHost[0]
	}
	translated, err := translate(text, from.String(), to.String(), false, 2, 0, nil)
	if err != nil {
		return "", err
	}

	return translated, nil
}

// TranslateWithParams translate a text with simple params as string
func TranslateWithParams(text string, params TranslationParams) (string, error) {
	if params.GoogleHost == "" {
		GoogleHost = "google.com"
	} else {
		GoogleHost = params.GoogleHost
	}
	translated, err := translate(text, params.From, params.To, true, params.Tries, params.Delay, params.Client)
	if err != nil {
		return "", err
	}
	return translated, nil
}
