package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type loggingTransport struct{}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s\n", reqDump)

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("RESPONSE:\n%s\n", respDump)

	return resp, nil
}

func main() {
	client := &http.Client{Transport: &loggingTransport{}}
	client.Get("https://httpbin.org/get")
}
