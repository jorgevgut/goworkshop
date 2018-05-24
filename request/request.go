package request

import (
	"net/http"
)

// Requester is the interface for perfoming requests
type Requester interface {
	Request(Options) chan *http.Response
}

// Options used to perform HTTP requests
type Options struct {
	URL     string
	Method  string
	Headers map[string]string
}
