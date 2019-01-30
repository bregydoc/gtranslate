package gtranslate

import (
	"fmt"
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
	{"Bye", "en", "es", "Adiós"},
	{"Hola", "es", "en", "Hello"},
	{"Adios", "es", "en", "Goodbye"},
	{"World", "en", "es", "Mundo"},
}

func TestTranslate(t *testing.T) {
	for i := 0; i < 5; i++ {
		for _, ta := range testingTable {
			translated, err := translate(ta.inText, ta.langFrom, ta.langTo, true)
			if err != nil {
				t.Error(err.Error())
			}
			if len(translated) < 2 {
				t.Fail()
			}
			fmt.Println("Translated", translated, "[OK]")
			if translated != ta.outText {
				t.Error("translated text is not the expected", ta.outText, " != ", translated)
			}
		}
	}
}

// TestGetGoogleTranslate is for testing propouse
// func TestGetGoogleTranslate(t *testing.T) {
// 	testText := "Some test text"
// 	for i := 0; i < 4; i++ {
// 		for _, ta := range testingTable {

// 			r, err := getGoogleTranslate(ta.inText, ta.langFrom, ta.langTo)
// 			if err != nil {
// 				t.Error(err.Error())
// 				t.Fail()
// 			}
// 			if r.StatusCode != http.StatusOK {

// 				t.Error("[" + strconv.Itoa(r.StatusCode) + "] failed request with text: '" + testText + "'")
// 			}
// 			if r.Body == nil {
// 				t.Fail()
// 			}
// 		}
// 	}

// }

// // TestRawTranslate testing rawTranslate function
// func TestRawTranslate(t *testing.T) {
// 	for i := 0; i < 4; i++ {
// 		for _, ta := range testingTable {
// 			data, err := rawTranslate(ta.inText, ta.langFrom, ta.langTo)
// 			if err != nil {
// 				t.Error(err.Error())
// 			}
// 			if len(data) < 10 {
// 				t.Fail()
// 			}
// 		}
// 	}

// }
