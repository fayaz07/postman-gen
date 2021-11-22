package core

import (
	"flag"
	"fmt"
)

const (
	usage = `
Hi! Thank you for choosing url generator from postman collection. 
Here is what you can do

Options:`
)

var version bool
var supportedLanguages bool

func InitFlags() {

	flag.BoolVar(&version, "version", false, "prints current version")
	flag.BoolVar(&supportedLanguages, "lang", false, "lists out supported languages")

	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}
}

func ParseFlags() {
	flag.Parse()

	if version {
		showVersion()
		return
	}

	if supportedLanguages {
		showSupportedLanguages()
		return
	}
}

func showVersion() {
	fmt.Println("v" + GetAppConfig().Version)
}

func showSupportedLanguages() {
	langs := GetAppConfig().Languages

	fmt.Println("This generator currently supports the following languages: ")

	for i := range langs {
		p := langs[i]
		fmt.Println("\t-", p) //p1, p2
	}
}
