package gtranslate

import (
	"net/http"
	"testing"
)

type testTable struct {
	inText   string
	langFrom string
	langTo   string
	outText  string
}

var testingTable = []testTable{
	{"Hello", "en", "es", "Hola"},
	{"Bye", "en", "es", "Adios"},
	{"Hola", "es", "en", "Hello"},
	{"Adios", "es", "en", "Bye"},
	{"World", "en", "es", "Mundo"},
}

// TestGetGoogleTranslate is for testing propouse
func TestGetGoogleTranslate(t *testing.T) {
	testText := "Some test text"
	for i := 0; i < 4; i++ {
		for _, ta := range testingTable {
			r, err := getGoogleTranslate(ta.inText, ta.langFrom, ta.langTo)
			if err != nil {
				t.Error(err.Error())
			}
			if r.StatusCode != http.StatusOK {
				t.Error("failed request with text: '" + testText)
			}
			if r.Body == nil {
				t.Fail()
			}
		}
	}

}

// TestRawTranslate testing rawTranslate function
func TestRawTranslate(t *testing.T) {
	for i := 0; i < 4; i++ {
		for _, ta := range testingTable {
			data, err := rawTranslate(ta.inText, ta.langFrom, ta.langTo)
			if err != nil {
				t.Error(err.Error())
			}
			if len(data) < 10 {
				t.Fail()
			}
		}
	}

}
