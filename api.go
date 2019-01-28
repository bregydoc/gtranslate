package gtranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/language"
)

const gt = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"

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

	raw, err := rawTranslate(text, from, to)
	if err != nil {
		return "", err
	}
	var resp []interface{}
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return "", err
	}

	responseText := ""
	for _, obj := range resp[0].([]interface{}) {

		if len(obj.([]interface{})) == 0 {
			break
		}
		responseText += obj.([]interface{})[0].(string)
	}

	return responseText, nil
}

func rawTranslate(text, from, to string) ([]byte, error) {
	r, err := getGoogleTranslate(text, from, to)

	if err != nil {
		if err == http.ErrHandlerTimeout {
			return nil, errBadNetwork
		}
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		// oh noooo
		return nil, errBadRequest
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getGoogleTranslate(text, from, to string) (*http.Response, error) {
	http.DefaultClient.Timeout = 10 * time.Second
	text = url.PathEscape(text)
	uri := fmt.Sprintf(gt, from, to, text)

	return http.Get(uri)
}
