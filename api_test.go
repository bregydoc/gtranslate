package gtranslate

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/text/language"
)

type testTable struct {
	inText   string
	langFrom interface{}
	langTo   interface{}
	outText  string
}

var testingTable = []testTable{
	{"Hello", "en", "es", "Hola"},
	{"Bye", language.English, "es", "Adiós"},
	{"Hola", "es", "en", "Hello"},
	{"Adios", "es", "en", "Goodbye"},
	{"World", language.English, "es", "Mundo"},
}

func TestTranslate(t *testing.T) {
	N := 5
	var totalDur time.Duration
	for i := 0; i < N; i++ {
		for _, ta := range testingTable {
			fromStr, toStr, err := fromToTypeConvert(ta.langFrom, ta.langTo)
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}

			start := time.Now()
			translated, err := translate(ta.inText, fromStr, toStr, true, 5, time.Second)
			if err != nil {
				t.Error(err.Error())
			}
			if len(translated) < 2 {
				t.Fail()
			}
			dur := time.Since(start)
			fmt.Print(".")
			totalDur += dur
			if translated != ta.outText {
				t.Error("translated text is not the expected", ta.outText, " != ", translated)
			}
		}
	}
	fmt.Println("\nMean time:", time.Duration(int(totalDur)/(len(testingTable)*N)))
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
