package template

import "github.com/majgis/htmls/token"

// Marshall HTML containing tokens into HTMLTemplate
func Marshall(name string, template []byte) (htmlTemplate HTMLTemplate, err error) {
	htmlTemplate = HTMLTemplate{Name: name}
	startPosition := 0
	endPosition := 0
	sectionStart := 0
	var tokens []token.HTMLToken
	var sections [][]byte

	// Iterate each character in template, split into sections surrounding tokens, marshall tokens
	for i, character := range template {
		switch character {

		// Start of Token
		case '{':
			if template[i+1] == '{' {
				startPosition = i + 2
			}

		// End of Token
		case '}':
			if template[i+1] == '}' {
				endPosition = i
				htmlToken, err := token.Marshall(template[startPosition:endPosition])
				if err != nil {
					break
				}
				tokens = append(tokens, htmlToken)
				sections = append(sections, template[sectionStart:startPosition-2])
				sectionStart = endPosition + 2
			}
		}
	}
	sections = append(sections, template[sectionStart:])
	htmlTemplate.Tokens = tokens
	htmlTemplate.Sections = sections
	return
}
