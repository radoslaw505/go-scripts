package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func postJSON() {
	data := map[string]string{"name": "John Doe", "email": "john@example.com"}
	jsonData, _ := json.Marshal(data)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("POST JSON status:", resp.Status)
}

func uploadFile() {
	file, _ := os.Open("file.txt")
	defer file.Close()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormFile("uploadfile", file.Name())
	io.Copy(fw, file)
	w.Close()

	req, _ := http.NewRequest("POST", "https://httpbin.org/post", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("Upload status:", resp.Status)
}

func main() {
	postJSON()
	uploadFile()
}
