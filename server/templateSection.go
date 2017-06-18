package server

import "github.com/majgis/htmls/token"

// A section of a HTMLTemplate used for final rendering
type templateSection struct {
	htmlToken token.HTMLToken
	bytes     []byte
	ch        chan responseChunk
}
