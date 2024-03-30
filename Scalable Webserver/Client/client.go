package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

func sendRequest(wg *sync.WaitGroup) {
    defer wg.Done() // Signal that this goroutine has finished

    _, err := http.Get("http://localhost:3000")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Sent request")
}

func main() {
    var wg sync.WaitGroup // Create an instance of WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1) // Increment the counter for each goroutine
        go sendRequest(&wg)
        time.Sleep(200 * time.Millisecond) // Control the rate of requests
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All requests sent")
}
