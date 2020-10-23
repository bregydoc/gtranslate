# gtranslate ![build](https://travis-ci.com/bregydoc/gtranslate.svg?branch=master)

Google Translate API for unlimited and free translations 📢.
This project was inspired by [google-translate-api](https://github.com/matheuss/google-translate-api) and [google-translate-token](https://github.com/matheuss/google-translate-token).

# Install

    go get github.com/bregydoc/gtranslate

# Use

```go
gtranslate.Translate("I'm alive", language.English, language.Spanish)
```

```go
gtranslate.Translate(text, en, es)
```

```go
gtranslate.Translate(text, language.English, es)
```

```go
gtranslate.TranslateA("I'm alive", language.Spanish)
```

```go
gtranslate.TranslateA(text, es)
```

```go
gtranslate.TranslateWithParams("I'm alive", gtranslate.TranslateWithParams{From: "en", To: "es"})
```


# Examples

## TranslateWithParams

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
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

## Translate

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "In the good old days of computing when memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc."
	translatedText, _ := gtranslate.Translate(text, en, es)

	fmt.Println("translated:", translatedText)
}
```

## Translate with language package

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func main() {
	text := "In the good old days of computing when memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc."
	translatedText, _ := gtranslate.Translate(text, language.English, language.Spanish)

	fmt.Println("translated:", translatedText)
}
```

## TranslateA

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "In the good old days of computing when memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc."
	translatedText, _ := gtranslate.Translate(text, es)

	fmt.Println("translated:", translatedText)
}
```

## TranslateA with language package

```go
package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func main() {
	text := "In the good old days of computing when memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc."
	translatedText, _ := gtranslate.TranslateA(text, language.Spanish)

	fmt.Println("translated:", translatedText)
}
```