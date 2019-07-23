package gtranslate

import (
	"testing"
)

func TestTranslateWithFromTo(t *testing.T) {
	for i := 0; i < 4; i++ {
		for _, ta := range testingTable {
			resp, err := TranslateWithParams(ta.inText, TranslationParams{
				From: ta.langFrom,
				To:   ta.langTo,
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
