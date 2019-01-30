package gtranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/text/language"

	"github.com/robertkrimen/otto"
)

var ttk otto.Value

func init() {
	ttk, _ = otto.ToValue("0")
}

func translate(text, from, to string, withVerification bool) (string, error) {
	if withVerification {
		if _, err := language.Parse(from); err != nil {
			fmt.Println("[WARNING], '" + from + "' is a invalid language, switching to 'auto'")
			from = "auto"
		}
		if _, err := language.Parse(to); err != nil {
			fmt.Println("[WARNING], '" + to + "' is a invalid language, switching to 'en'")
			to = "en"
		}
	}
	t, _ := otto.ToValue(text)
	urll := "https://translate.google.com/translate_a/single"
	token := get(t, ttk)

	data := map[string]string{
		"client": "gtx",
		"sl":     from,
		"tl":     to,
		"hl":     to,
		// "dt":     []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"},
		"ie":   "UTF-8",
		"oe":   "UTF-8",
		"otf":  "1",
		"ssel": "0",
		"tsel": "0",
		"kc":   "7",
		"q":    text,
	}
	u, err := url.Parse(urll)
	if err != nil {
		return "", nil
	}
	parameters := url.Values{}
	for k, v := range data {
		parameters.Add(k, v)
	}
	for _, v := range []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"} {
		parameters.Add("dt", v)
	}
	parameters.Add(token["name"], token["value"])
	u.RawQuery = parameters.Encode()
	// fmt.Println(u)
	r, err := http.Get(u.String())

	if err != nil {
		if err == http.ErrHandlerTimeout {
			return "", errBadNetwork
		}
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		// oh noooo
		return "", errBadRequest
	}

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var resp []interface{}
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return "", err
	}

	// pp.Println(resp)

	responseText := ""
	for _, obj := range resp[0].([]interface{}) {

		if len(obj.([]interface{})) == 0 {
			break
		}
		responseText += obj.([]interface{})[0].(string)
	}

	return responseText, nil
}
