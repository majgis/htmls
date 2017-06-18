package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/majgis/htmls/template"
)

// Start will setup routes based on template names and start server listening on given port
func Start(port int, htmlTemplates []template.HTMLTemplate) {

	// Handle each template using name as route
	for _, htmlTemplate := range htmlTemplates {
		http.HandleFunc("/"+htmlTemplate.Name, templateHandlerFactory(htmlTemplate))
	}

	// Handle root to provide list of routes
	http.HandleFunc("/", rootHandlerFactory(htmlTemplates))

	// Start the server
	fmt.Println("Starting server on " + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}
