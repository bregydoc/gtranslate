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
