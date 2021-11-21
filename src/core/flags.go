package core

import (
	"flag"
	"fmt"
	"os"
)

const (
	usage = `usage: %s
	Generate url's file from a postman collection
	Options:
`
)

func InitFlags() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func showVersion() {
	fmt.Println()
}
