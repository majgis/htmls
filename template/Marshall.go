package template

import "github.com/majgis/htmls/token"

// Marshall HTML containing tokens into HTMLTemplate
func Marshall(name string, template []byte) (htmlTemplate HTMLTemplate, err error) {
	htmlTemplate = HTMLTemplate{Name: name}
	startCount := 0
	startPosition := 0
	endCount := 0
	endPosition := 0
	sectionStart := 0
	var tokens []token.HTMLToken
	var sections [][]byte

	// Iterate each character in template, split into sections surrounding tokens, marshall tokens
	for i, character := range template {
		switch character {

		// Start of Token
		case '{':
			startCount++
			if startCount == 2 {
				startCount = 0
				startPosition = i + 1
			}

		// End of Token
		case '}':
			endCount++
			if endCount == 2 {
				endCount = 0
				endPosition = i - 2
				htmlToken, err := token.Marshall(template[startPosition : endPosition+1])
				if err != nil {
					break
				}
				tokens = append(tokens, htmlToken)
				sections = append(sections, template[sectionStart:startPosition-2])
				sectionStart = i + 1
			}
		}
	}
	sections = append(sections, template[sectionStart:])
	htmlTemplate.Tokens = tokens
	htmlTemplate.Sections = sections
	return
}
