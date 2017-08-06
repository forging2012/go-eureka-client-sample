package eureka

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"
)

// Accepts a Httpaction and a one-way channel to write the results to.
func DoHttpRequest(httpAction HttpAction) bool {
	req := buildHttpRequest(httpAction)

	var DefaultTransport http.RoundTripper = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp, err := DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Printf("HTTP request failed: %s", err)
		return false
	} else {
		return true
		defer resp.Body.Close()
	}
	return false
}

func buildHttpRequest(httpAction HttpAction) *http.Request {
	var req *http.Request
	var err error
	if httpAction.Body != "" {
		reader := strings.NewReader(httpAction.Body)
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
	} else if httpAction.Template != "" {
		reader := strings.NewReader(httpAction.Template)
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
	} else {
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, nil)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Add headers
	for name, value := range httpAction.Headers {
		req.Header.Add(name, value)
	}
	return req
}
