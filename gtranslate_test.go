package gtranslate

import (
	"testing"
)

func TestTranslateWithFromTo(t *testing.T) {
	for i := 0; i < 10; i++ {
		for _, ta := range testingTable {
			resp, err := TranslateWithFromTo(ta.inText, FromTo{
				From: ta.langFrom,
				To:   ta.langTo,
			})
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}

			if resp != ta.outText {
				t.Error("translated text is not the expected translate")
			}
		}
	}
}
