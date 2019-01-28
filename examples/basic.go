package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "a"
	translated, err := gtranslate.TranslateWithFromTo(
		text,
		gtranslate.FromTo{
			From: "en",
			To:   "es",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | es: %s \n", text, translated)
	// en: Hello World | es: Hola Mundo
}
