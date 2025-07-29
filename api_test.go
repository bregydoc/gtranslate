package gtranslate

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
	{"Adios", "es", "en", "Bye"},
	{"World", "en", "es", "Mundo"},
}

func TestTranslate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[[["Hola","Hello","",""]],,"es",,,[["Hello",
		[5],true,false]],[["Hello",5,true,false]],,,[]]`))
	}))
	defer server.Close()

	for _, ta := range testingTable {
		GoogleHost = server.URL
		translated, err := translate(ta.inText, ta.langFrom, ta.langTo, true, 1, 0)
		assert.NoError(t, err)
		assert.Equal(t, ta.outText, translated)
	}
}
