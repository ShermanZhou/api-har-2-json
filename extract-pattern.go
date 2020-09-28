package main

import (
	"fmt"
	"regexp"
)

var prefix string = "/api/v1/"

// the capture group will be used for export file base name
var patterns = []ExtractPattern{
	/*{
		Verb:     "GET",
		Url:      `(accounts/\d+)$`,
		StatusOK: true,
	},
	{
		Verb:     "GET",
		Url:      `(clients/\d+/accounts)$`,
		StatusOK: true,
	},
	{
		Verb:     "GET",
		Url:      `(clients/\d+)$`,
		StatusOK: true,
	},
	{
		Verb:     "*",
		Url:      `(applications/\d+)$`,
		StatusOK: true,
	},
	{
		Verb:     "GET",
		Url:      `(applications/\d+/related)$`,
		StatusOK: true,
	},*/
	{
		Verb:     "*",
		Url:      `(.*)`,
		StatusOK: true,
	},
}

type ExtractPattern struct {
	Verb     string // GET, PUT, POST, *
	Url      string
	StatusOK bool // 200 serials
	RegEx    *regexp.Regexp
}

func compilePatterns(patterns []ExtractPattern) {
	for index, p := range patterns {
		reg, err := regexp.Compile(prefix + p.Url)
		if err != nil {
			fmt.Printf("pattern is invalid: %s\n", p.Url)
		}
		patterns[index].RegEx = reg
	}
}
