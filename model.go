package main

import (
	"regexp"
)

type Har struct {
	Log Log `json:"log"`
}

type Log struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Request struct {
	Url      string   `json:"url"`
	Headers  []Header `json:'headers'`
	Method   string   `json:"method"`
	PostData PostData `json:"postData"`
}
type PostData struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}

type Response struct {
	Content    Content  `json:"content"`
	Headers    []Header `json:"headers"`
	Status     int      `json:"status"`
	StatusText string   `json:"statusText"`
}

type Content struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}

type MatchedEntry struct {
	Entry Entry
	RegEx *regexp.Regexp
}
