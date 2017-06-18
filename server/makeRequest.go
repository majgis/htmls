package server

import (
	"io/ioutil"
	"net/http"
)

// Make a request to the given url with chunks of response passed to the given channel.
// The channel will be closed on completion, errors besides EOF are written to the channel.
// url: The address to make the request to
// ch: The channel to write response chunks back to
func makeRequest(url string, timeoutMS int, ch chan<- responseChunk) {
	response, err := http.Get(url)
	responseChunk := responseChunk{}
	// If an error occurred, write error to channel, close channel and exit
	if err != nil {
		responseChunk.Error = err
		ch <- responseChunk
		close(ch)
		return
	}
	responseChunk.Bytes, responseChunk.Error = ioutil.ReadAll(response.Body)
	ch <- responseChunk
	close(ch)
	return
}
