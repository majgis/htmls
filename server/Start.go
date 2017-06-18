package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/majgis/htmls/template"
)

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