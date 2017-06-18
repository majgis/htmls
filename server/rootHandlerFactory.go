package server

import (
	"fmt"
	"net/http"

	"github.com/majgis/htmls/template"
)

// rootHandlerFactory returns function to serve page with list of routes
func rootHandlerFactory(htmlTemplates []template.HTMLTemplate) func(http.ResponseWriter, *http.Request) {
	rootTemplate := []byte("<h3>Routes:</h3><ul>")
	for _, htmlTemplate := range htmlTemplates {
		line := fmt.Sprintf("<li><a href=\"/%s\">%s</a></li>", htmlTemplate.Name, htmlTemplate.Name)
		rootTemplate = append(rootTemplate, line...)
	}
	rootTemplate = append(rootTemplate, "</ul>"...)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(rootTemplate)
		return
	}
}
