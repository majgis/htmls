package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/majgis/htmls/server"
	"github.com/majgis/htmls/template"
)

const usage = `
  Usage:
    htmls template.html
`

func main() {

	// Print usage and exit if an argument is not given
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	// Read the file and exit if there is an error
	templatePath := os.Args[1]
	contents, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Println("The template file could not be read: " + templatePath)
		os.Exit(1)
	}

	// Marshall the template and exit if there is an error
	htmlTemplate, err := template.Marshall(contents)
	if err != nil {
		fmt.Println("The template is malformed: " + err.Error())
		os.Exit(1)
	}

	// Start the server
	htmlTemplates := []template.HTMLTemplate{htmlTemplate}
	server.Start(htmlTemplates)
}
