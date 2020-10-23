package gtranslate

import (
	"time"
)

var GoogleHost = "google.com"

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
}

// Translate translate a text using native tags offer by go language
func Translate(text string, from, to interface{}, googleHost ...string) (string, error) {
	fromStr, toStr, err := fromToTypeConvert(from, to)
	if err != nil {
		return "", err
	}

	if len(googleHost) != 0 && googleHost[0] != "" {
		GoogleHost = googleHost[0]
	}

	translated, err := translate(text, fromStr, toStr, false, defaultNumberOfRetries, 0)
	if err != nil {
		return "", err
	}

	return translated, nil
}

// TranslateA translate a text using native tags offer by go language
func TranslateA(text string, to interface{}, googleHost ...string) (string, error) {
	fromStr, toStr, err := fromToTypeConvert("auto", to)
	if err != nil {
		return "", err
	}

	if len(googleHost) != 0 && googleHost[0] != "" {
		GoogleHost = googleHost[0]
	}
	translated, err := translate(text, fromStr, toStr, false, defaultNumberOfRetries, 0)
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
	translated, err := translate(text, params.From, params.To, true, params.Tries, params.Delay)
	if err != nil {
		return "", err
	}
	return translated, nil
}
