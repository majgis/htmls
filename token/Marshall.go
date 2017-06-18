package token

import (
	"errors"
	"regexp"
	"strconv"
)

// Marshall extracts information from byte slice containing HTML token fields
// uri: Address of HTML endpoint
// cache_ms: Duration in milliseconds to retain content before issuing a new request
// timeout_ms: Duration in milliseconds to wait before canceling a request
func Marshall(token []byte) (htmlToken HTMLToken, err error) {
	uriRegexp := regexp.MustCompile(`uri="(.*?)"`)
	cacheRegexp := regexp.MustCompile(`cache_ms="(.*?)"`)
	timeoutRegexp := regexp.MustCompile(`timeout_ms="(.*?)"`)
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
