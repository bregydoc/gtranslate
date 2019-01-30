package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

func main1() {
	// a, err := otto.ToValue("hello")
	// if err != nil {
	// 	panic(err)
	// }
	// err = gtranslate.GetSM(a)
	// if err != nil {
	// 	panic(err)
	// }
	// text, _ := otto.ToValue("Hello")
	// v, _ := otto.ToValue("0")
	// // gtranslate.UpdateTTK(v)
	// fmt.Println(gtranslate.Get(text, v))
	t, _ := gtranslate.Translate(`
	In the good old days of computing when memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc.
	`, language.English, language.Spanish)

	fmt.Println("translated:", t)
}
