# gtranslate ![build](https://travis-ci.com/bregydoc/gtranslate.svg?branch=master)

Google translate API for unlimited and free translations.
This project was inspired by [google-translate-api](https://github.com/matheuss/google-translate-api).

# Install

    go get github.com/bregydoc/gtranslate

# Use

```go
gtranslate.Translate("I'm alive", language.English, language.Spanish)
```

```go
gtranslate.TranslateWithFromTo("I'm alive", gtranslate.FromTo{From: "en", To: "es"})
```

# Example

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithFromTo(
		text,
		gtranslate.FromTo{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}
```
