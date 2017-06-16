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
	fmt.Println(string(contents))
}
