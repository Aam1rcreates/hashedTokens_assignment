# Simple Scalable Webserver

## Problem
Build a simple web server that scales as it gets more requests. The idea is to spawn more
goroutines as the same endpoint is receiving lots of requests.
- Put a sleep operation to block a specific goroutine for a finite duration of time. This
simulates a compute-heavy operation
- Have only one endpoint in your application
- Have a client that sends a lot of requests within the sleep interval.

## Solution

### Idea
We want to create a simple web server that can handle increasing loads by spawning more goroutines 
(lightweight threads) as the same endpoint receives lots of requests. The server should have the 
following characteristics:

- Single Endpoint: The server will have only one endpoint.
- Compute-Heavy Operation Simulation: To simulate a compute-heavy operation, we’ll introduce a
 sleep operation within the request handler.
- Scalability: As the load increases, the server should spawn more goroutines to handle
 incoming requests.

### Approach

#### 1. Initialize the Web Server:
- Create a Go program that sets up an HTTP server listening on a specific port (e.g., 8080).
- Define a single endpoint (e.g., “/”) that will handle all incoming requests.
#### 2. Request Handler Function with Channels:
- Implement a function (let’s call it handleRequests) that will be called for each
- incoming request.
- Inside this function:
  - Create a channel (e.g., requestChan) to communicate between the main goroutine and
  - request-handling goroutines.
  - Send the request to the channel.
  - Spawn a new goroutine to handle the request.
  - In the spawned goroutine, simulate a compute-heavy operation by introducing a
  - sleep (e.g., 2 seconds).
  - After the operation, send a response (e.g., a success message) back to the main
  - goroutine via another channel (e.g., responseChan).
#### 3. Goroutine Spawning:
  - In the main goroutine (server), listen for incoming requests.
  - For each request, spawn a new goroutine to handle it.
  - Wait for the response from the request-handling goroutine via the responseChan.
  - Send the response to the client.

#### 4. Client for Load Testing:
- Create a separate Go program (the client) that sends a large number of requests to
- the server.
- The client should:
  - Make HTTP GET requests to the server’s endpoint.
  - Sleep for a short duration (e.g., 200 milliseconds) between each request to control the
    rate of requests.

### Pseudo Code

```
// main.go (Server)

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequests(requestChan chan<- struct{}, responseChan <-chan string) {
	for {
		// Wait for an incoming request
		<-requestChan

		// Simulate a compute-heavy operation
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

		fmt.Fprintf(w, response)
	})

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

```
