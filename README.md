# wild

## golang wildcard matching package

A golang adaptation of the FastWildComparePortable wildcard matching alogoritm created by Kirk J Krauss.

The first incarnation of the algorithm was documented in the[ Dr Dobbs journal August 26, 2008](https://www.drdobbs.com/architecture-and-design/matching-wildcards-an-algorithm/210200888) and updated in [October 7, 2014](https://www.drdobbs.com/architecture-and-design/matching-wildcards-an-empirical-way-to-t/240169123).

The further refinements were documented on [http://developforperformance.com](http://developforperformance.com) website in the article ["Matching Wildcards: An Improved Algorithm for Big Data"](http://developforperformance.com/MatchingWildcards_AnImprovedAlgorithmForBigData.html) and the accompanying source code was published on [https://github.com/kirkjkrauss/MatchingWildcards](https://github.com/kirkjkrauss/MatchingWildcards) under a Apache v2 license.

The golang adoptation is based on the pointerless implementation [FastWildComparePortable](https://github.com/kirkjkrauss/MatchingWildcards/blob/master/Listing2.cpp). As the FastWildComparePortable implementation depends on null terminitated strings, this package turns the input strings in to an array of runes, and appends a '\x00' rune, to enable an one to one translation of the logic flow of the algorithm. 

This decision will be revisted when needed, but will have no impact on the public signature of the package and its exposed functions.



### Installation

    go get -u github.com/gertd/wild

### Example usage

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
		flag.BoolVar(&caseinsensitive, "caseinsensitive", false, "caseinsensitive comparison")
		flag.Parse()
	
		match := wild.Match(pattern, search, caseinsensitive)
		fmt.Printf("pattern: [%s] search: [%s] matched: %t\n", pattern, search, match)
	}

### Results

	> go run ./cmd/wc/main.go  --pattern "ab*" --search "Abc"
	pattern: [ab*] search: [Abc] matched: false
	
	> go run ./cmd/wc/main.go  --pattern "ab*" --search "Abc" --caseinsensitive
	pattern: [ab*] search: [Abc] matched: true