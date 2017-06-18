package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/majgis/htmls/template"
	"github.com/majgis/htmls/token"
)

var htmlErrorComment = []byte("<!-- ERROR -->")

// Data passed through channel from requested response
type responseChunk struct {
	Bytes []byte
	Error error
}

// Make a request to the given url with chunks of response passed to the given channel.
// The channel will be closed on completion, errors besides EOF are written to the channel.
// url: The address to make the request to
// ch: The channel to write response chunks back to
func makeRequest(url string, timeoutMS int, ch chan<- responseChunk) {
	response, err := http.Get(url)
	responseChunk := responseChunk{}
	// If an error occurred, write error to channel, close channel and exit
	if err != nil {
		responseChunk.Error = err
		ch <- responseChunk
		close(ch)
		return
	}
	responseChunk.Bytes, responseChunk.Error = ioutil.ReadAll(response.Body)
	ch <- responseChunk
	close(ch)
	return
}

// A section of a HTMLTemplate
type templateSection struct {
	htmlToken token.HTMLToken
	bytes     []byte
	ch        chan responseChunk
}

// Generate a function with htmlTemplate stored in closure
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

// Start will setup routes based on template names and start server listening on given port
func Start(port int, htmlTemplates []template.HTMLTemplate) {

	// Server will handle each template using name as route
	for _, htmlTemplate := range htmlTemplates {
		http.HandleFunc("/"+htmlTemplate.Name, handlerFactory(htmlTemplate))
	}

	// Start the server
	fmt.Println("Starting server on " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
