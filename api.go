package gtranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/language"
)

const gt = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dj=1&dt=t&ie=UTF-8&q=%s"

func translateOld(text, from, to string, withVerification bool) (string, error) {
	if withVerification {
		if _, err := language.Parse(from); err != nil && from != "auto" {
			log.Println("[WARNING], '" + from + "' is a invalid language, switching to 'auto'")
			from = "auto"
		}
		if _, err := language.Parse(to); err != nil {
			log.Println("[WARNING], '" + to + "' is a invalid language, switching to 'en'")
			to = "en"
		}
	}

	raw, err := rawTranslate(text, from, to)
	if err != nil {
		return "", err
	}
	resp := new(struct {
		Sentences []struct {
			Trans   string `json:"trans"`
			Orig    string `json:"orig"`
			Backend int    `json:"backend"`
		} `json:"sentences"`
		Src   string `json:"src"`
		Spell struct {
		} `json:"spell"`
	})
	err = json.Unmarshal(raw, resp)
	if err != nil {
		return "", err
	}

	responseText := ""
	responseText = resp.Sentences[0].Trans

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

func getGoogleTranslate(text, from, to string, customClient ...*http.Client) (*http.Response, error) {
	client := http.DefaultClient
	if len(customClient) != 0 {
		client = customClient[0]
		log.Println("[WARNING] Using custom proxy")
	}

	client.Timeout = 40 * time.Second

	text = url.PathEscape(text)
	uri := fmt.Sprintf(gt, from, to, text)

	return client.Get(uri)
}
