package gtranslate

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateWithFromTo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[[["Hola","Hello","",""]],,"es",,,[["Hello",
		[5],true,false]],[["Hello",5,true,false]],,,[]]`))
	}))
	defer server.Close()

	for _, ta := range testingTable {
		resp, err := TranslateWithParams(ta.inText, TranslationParams{
			From:       ta.langFrom,
			To:         ta.langTo,
			Tries:      1,
			GoogleHost: server.URL,
		})
		assert.NoError(t, err)
		assert.Equal(t, ta.outText, resp)
	}
}
