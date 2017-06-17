package server

import (
	"fmt"

	"github.com/majgis/htmls/template"
)

func Start(htmlTemplates []template.HTMLTemplate) {

	// Print the URI for every token
	for _, htmlTemplate := range htmlTemplates {
		for _, token := range htmlTemplate.Tokens {
			fmt.Println("URI:", token.URI)
		}
	}
}