package server

import (
	"net/http"
)

// Make a request to the given url with chunks of response passed to the given channel.
// The channel will be closed on completion, errors besides EOF are written to the channel.
// url: The address to make the request to
// ch: The channel to write response chunks back to
func makeRequest(url string, timeoutMS int, ch chan<- responseChunk) {
	response, err := http.Get(url)

	// If an error occurred, write error to channel, close channel and exit
	if err != nil {
		ch <- responseChunk{Error: err}
		close(ch)
		return
	}

	// Iterate through chunks of the response body
	for {
		chunk := responseChunk{Bytes: make([]byte, 2000, 2000)}
		byteCount, err := response.Body.Read(chunk.Bytes)
		ch <- chunk
		if byteCount == 0 || chunk.Error != nil {
			chunk.Error = err
			break
		}
	}
	close(ch)

	return
}
