package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request timed out:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Observability example status:", resp.Status)
}
