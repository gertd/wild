package main

import (
	"fmt"

	"github.com/gertd/wild"
	flag "github.com/spf13/pflag"
)

func main() {
	var (
		pattern         string
		search          string
		caseinsensitive bool
	)

	flag.StringVar(&pattern, "pattern", "", "wildcard pattern")
	flag.StringVar(&search, "search", "", "searched content string")
	flag.BoolVar(&caseinsensitive, "case-insensitive", false, "case-insensitive comparison")
	flag.Parse()

	match := wild.Match(pattern, search, caseinsensitive)
	fmt.Printf("pattern: [%s] search: [%s] matched: %t\n", pattern, search, match)
}
