package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	out, _ := os.Create("largefile.bin")
	defer out.Close()

	resp, err := http.Get("https://speed.hetzner.de/100MB.bin")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)
	fmt.Println("Downloaded large file")
}
