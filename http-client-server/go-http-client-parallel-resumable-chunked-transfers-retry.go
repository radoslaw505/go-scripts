package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// doWithRetry retries a function up to maxRetries times with exponential backoff.
func doWithRetry(fn func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		wait := time.Duration(1<<i) * time.Second
		fmt.Printf("Retry %d after error: %v (waiting %v)\n", i+1, err, wait)
		time.Sleep(wait)
	}
	return err
}

// downloadChunk downloads a specific byte range from the URL and writes it to the file at the correct offset.
func downloadChunk(url string, start, end int64, file *os.File) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	return doWithRetry(func() error {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status: %s", resp.Status)
		}

		// Seek to the correct position and write the chunk
		_, err = file.Seek(start, io.SeekStart)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, resp.Body)
		return err
	}, 3)
}

func main() {
	url := "https://speed.hetzner.de/100MB.bin" // Example large file
	output := "parallel_download.bin"

	// Known file size (could be fetched via HEAD request)
	fileSize := int64(100 * 1024 * 1024) // 100 MB
	chunkSize := int64(5 * 1024 * 1024)  // 5 MB

	// Create output file
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	for start := int64(0); start < fileSize; start += chunkSize {
		end := start + chunkSize - 1
		if end >= fileSize {
			end = fileSize - 1
		}
		wg.Add(1)
		go func(s, e int64) {
			defer wg.Done()
			fmt.Printf("Downloading bytes %d-%d\n", s, e)
			if err := downloadChunk(url, s, e, file); err != nil {
				fmt.Printf("Failed chunk %d-%d: %v\n", s, e, err)
			} else {
				fmt.Printf("Completed chunk %d-%d\n", s, e)
			}
		}(start, end)
	}

	wg.Wait()
	fmt.Println("Download complete:", output)
}
