package server

// Data passed through channel from requested response
type responseChunk struct {
	Bytes []byte
	Error error
}
