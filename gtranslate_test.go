package gtranslate

import (
	"strings"
	"testing"
	"time"
)

func TestTranslateWithFromTo(t *testing.T) {
	for i := 0; i < 4; i++ {
		for _, ta := range testingTable {
			fromStr, toStr, err := fromToTypeConvert(ta.langFrom, ta.langTo)
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}

			resp, err := TranslateWithParams(ta.inText, TranslationParams{
				From:       fromStr,
				To:         toStr,
				Tries:      5,
				Delay:      time.Second,
				GoogleHost: "google.cn",
			})
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}
			if resp != ta.outText {
				t.Error("translated text is not the expected", ta.outText, " != ", resp)
			}
		}
	}
}

func TestTranslateA(t *testing.T) {
	for _, ta := range testingTable {
		translated, err := TranslateA(ta.inText, ta.langTo)
		if err != nil {
			t.Error(err, err.Error())
			t.Fail()
		}

		ta.outText, translated = strings.ToLower(ta.outText), strings.ToLower(translated)

		if ta.outText != translated {
			t.Errorf("TestTranslateA(%s,%v) = %q. Want %q", ta.inText, ta.inText, translated, ta.outText)
		}
	}

}

func TestTranslateFT(t *testing.T) {
	for _, ta := range testingTable {
		translated, err := Translate(ta.inText, ta.langFrom, ta.langTo)
		if err != nil {
			t.Error(err, err.Error())
			t.Fail()
		}

		ta.outText, translated = strings.ToLower(ta.outText), strings.ToLower(translated)

		if ta.outText != translated {
			t.Errorf("TestTranslateA(%s,%v) = %q. Want %q", ta.inText, ta.inText, translated, ta.outText)
		}
	}

}
