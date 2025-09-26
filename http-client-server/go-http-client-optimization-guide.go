package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}
	client := &http.Client{Transport: tr}

	resp, _ := client.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	fmt.Println("Connection pooling status:", resp.Status)
}
