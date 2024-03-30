package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequests(requestChan chan struct{}, responseChan chan string) {
	for {
		// Wait for an incoming request
		<-requestChan

		// 2 seconds of sleep to simulate compute-heavy operation
		time.Sleep(2 * time.Second)

		// Send response back to main goroutine
		responseChan <- "Request handled successfully!"
	}
}

func main() {
	requestChan := make(chan struct{})
	responseChan := make(chan string)

	// Start request handler goroutine
	go handleRequests(requestChan, responseChan)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send request to requestChan
		requestChan <- struct{}{}

		// Wait for response from responseChan
		response := <-responseChan

		fmt.Println(response)
	})

	fmt.Println("Server started on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
