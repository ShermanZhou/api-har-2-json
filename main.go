package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var logInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC)
var logErr = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.LUTC)

func main() {
	var harPath = flag.String("har", "", "har file path")
	flag.Parse()
	if *harPath == "" {
		panic("no harfile given in parameter, please run --help")
	}
	var fileName string
	if !filepath.IsAbs(*harPath) {
		fileName = filepath.Join(currentExecutablePath(), *harPath)
	} else {
		fileName = *harPath
	}
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var har Har
	err = json.Unmarshal(raw, &har)
	if err != nil {
		panic(err)
	}

	compilePatterns(patterns)
	// har file api can be called more than once, last will win
	// matchedEntries := make(map[string]MatchedEntry)
	var entryCounter = 0
	for _, entry := range har.Log.Entries {
		if strings.Contains(entry.Request.Url, "/api/") {
			for _, p := range patterns {
				if p.RegEx == nil {
					continue
				}
				if p.RegEx.MatchString(entry.Request.Url) {
					entryURL := entry.Request.Url
					if (p.Verb != "*" && p.Verb == entry.Request.Method) || p.Verb == "*" {
						logInfo.Printf("Rule match for %s, extract \n", entryURL)
						entryCounter++
						fitURL := p.RegEx.FindStringSubmatch(entry.Request.Url)
						fileBaseName := makeExtractFileName(fitURL[1], entry.Request.Method, entryCounter)
						var raw []byte
						if entry.Request.Method != "GET" {
							raw = []byte(entry.Request.PostData.Text)
						} else {
							raw = []byte(entry.Response.Content.Text)
						}
						fileBaseName = fileBaseName + ".json"
						logInfo.Printf("export %d: %s\n", entryCounter, fileBaseName)
						err := ioutil.WriteFile(fileBaseName, raw, 0666)
						if err != nil {
							logErr.Printf("file creation failed: %s %s \n", fitURL, err.Error())
						}
					}
				}
			}
		}
	}
}

func makeExtractFileName(url string, verb string, seq int) string {
	return fmt.Sprintf("%d-%s-[%s]", seq, strings.ReplaceAll(url, "/", "-"), verb)
}

func currentExecutablePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logErr.Panicln(err.Error())
	}
	return dir
}
