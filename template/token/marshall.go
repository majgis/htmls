package token

import (
	"errors"
	"regexp"
	"strconv"
)

var uriRegexp = regexp.MustCompile(`uri="(.*?)"`)
var cacheRegexp = regexp.MustCompile(`cache_ms="(.*?)"`)
var timeoutRegexp = regexp.MustCompile(`timeout_ms="(.*?)"`)

// HTMLToken is an in-memory object for the token within an HTML Template
// URI: The URI to the endpoint that will provide HTML snippet
// CacheMS: The duration in milliseconds that the HTML snippet should be stored
type HTMLToken struct {
	URI       string
	CacheMS   int
	TimeoutMS int
}

// Marshall html token to HTMLToken
func Marshall(token []byte) (htmlToken HTMLToken, err error) {
	htmlToken = HTMLToken{}
	tokenString := string(token)

	// The uri field is required
	uriMatch := uriRegexp.FindStringSubmatch(tokenString)
	if len(uriMatch) < 1 {
		err = errors.New("A uri field is required for all tokens")
		return
	}
	htmlToken.URI = uriMatch[1]

	// The cache_ms field is optional
	cacheMatch := cacheRegexp.FindStringSubmatch(tokenString)
	if len(cacheMatch) > 0 {
		htmlToken.CacheMS, err = strconv.Atoi(cacheMatch[1])
		if err != nil {
			return
		}
	}

	// The timeout_ms field is optional, default is 1000
	timeoutMatch := timeoutRegexp.FindStringSubmatch(tokenString)
	if len(timeoutMatch) > 0 {
		htmlToken.TimeoutMS, err = strconv.Atoi(timeoutMatch[1])
		if err != nil {
			return
		}
	} else {
		htmlToken.TimeoutMS = 1000
	}

	return
}
