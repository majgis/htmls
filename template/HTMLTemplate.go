package template

import "github.com/majgis/htmls/token"

// HTMLTemplate is an in-memory representation of html template
// Name: The name of the template taken from HTML filename
// Sections: HTML sections split by the token
// Tokens: HTMLToken objects holding configuration
type HTMLTemplate struct {
	Name     string
	Sections [][]byte
	Tokens   []token.HTMLToken
}
