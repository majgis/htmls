package template

// HTMLTemplate is an in-memory representation of html template
type HTMLTemplate struct {
	Sections [][]byte
	Tokens   [][]byte
}

// Marshall template into htmlTemplate
func Marshall(template []byte) (HTMLTemplate, error) {
	t := HTMLTemplate{}
	startCount := 0
	startPosition := 0
	endCount := 0
	endPosition := 0
	var tokens [][]byte
	for i, v := range template {
		if v == '{' {
			startCount++
			if startCount == 2 {
				startCount = 0
				startPosition = i + 1
			}
		} else if v == '}' {
			endCount++
			if endCount == 2 {
				endCount = 0
				endPosition = i - 2
				tokens = append(tokens, template[startPosition:endPosition+1])
			}
		}
	}
	t.Tokens = tokens
	return t, nil
}
