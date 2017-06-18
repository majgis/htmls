package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/majgis/htmls/template"
)

var htmlErrorComment = []byte("<!-- ERROR -->")

// Generate a function for handling requests to a given HTML template
func handlerFactory(htmlTemplate template.HTMLTemplate) func(http.ResponseWriter, *http.Request) {
	firstSectionBytes := htmlTemplate.Sections[0]
	otherSectionsBytes := htmlTemplate.Sections[1:]

	return func(w http.ResponseWriter, r *http.Request) {

		// marshall sections and make concurrent requests
		templateSections := []templateSection{}
		for i, sectionBytes := range otherSectionsBytes {
			templateSection := templateSection{
				htmlToken: htmlTemplate.Tokens[i],
				bytes:     sectionBytes,
				ch:        make(chan responseChunk, 1),
			}
			templateSections = append(templateSections, templateSection)
			go makeRequest(templateSection.htmlToken.URI, templateSection.htmlToken.TimeoutMS, templateSection.ch)
		}

		// Write first section to response
		_, err := w.Write(firstSectionBytes)
		if err != nil {
			return
		}

		// Iterate through sections
		for _, templateSection := range templateSections {

			// Iterate through channel
			for responseChunk := range templateSection.ch {

				// If error, write html error comment and break
				if responseChunk.Error != nil {
					_, err = w.Write(htmlErrorComment)
					break
				}

				// Write chunk to response
				_, err = w.Write(responseChunk.Bytes)
				fmt.Println("here", strconv.Itoa(len(responseChunk.Bytes)))
				if err != nil {
					break
				}
			}
			if err != nil {
				break
			}

			// Write remainder of section
			_, err = w.Write(templateSection.bytes)
			if err != nil {
				break
			}

		}

		return
	}
}
