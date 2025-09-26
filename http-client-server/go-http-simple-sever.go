package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Configure Endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server!")
	})

	// Configure Server
	// const serverAddr string = "127.0.0.1:3000"
	const port string = ":8080"

	// fmt.Println("Server listening on port 3000")
	fmt.Println("Server listening on port: ", port)
	// err := http.ListenAndServe(serverAddr, nil)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("error starting server", err)
	}
}
