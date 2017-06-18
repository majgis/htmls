package token

// HTMLToken is an in-memory object for the token within an HTML Template
// URI: The URI to the endpoint that will provide HTML snippet
// CacheMS: The duration in milliseconds that the HTML snippet should be stored
// TimeoutMS: The duration in milliseconds to wait before timing out
type HTMLToken struct {
	URI       string
	CacheMS   int
	TimeoutMS int
}
