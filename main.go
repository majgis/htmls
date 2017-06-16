package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const usage = `
  Usage:
    htmls template.html
`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	templatePath := os.Args[1]
	contents, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Println("The template file could not be read: " + templatePath)
		os.Exit(1)
	}
	template, err := marshallTemplate(contents)
	if err != nil {
		fmt.Println("The template is malformed: " + err.Error())
		os.Exit(1)
	}
	for _, v := range template.tokens {
		fmt.Println(string(v))
	}
}

type htmlTemplate struct {
	sections [][]byte
	tokens   [][]byte
}

// Marshall template into htmlTemplate
func marshallTemplate(template []byte) (htmlTemplate, error) {
	t := htmlTemplate{}
	startCount := 0
	startPosition := 0
	endCount := 0
	endPosition := 0
	var tokens [][]byte
	for i, v := range template {
		if v == '{' {
			startCount++
			if startCount == 2 {
				startCount = 0
				startPosition = i + 1
			}
		} else if v == '}' {
			endCount++
			if endCount == 2 {
				endCount = 0
				endPosition = i - 2
				tokens = append(tokens, template[startPosition:endPosition+1])
			}
		}
	}
	t.tokens = tokens
	return t, nil
}
