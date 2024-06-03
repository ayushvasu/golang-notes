package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	fmt.Println("starting !!")

	c := make([]chan string, len(links)) // Create slice of channels with length equal to links

	var wg sync.WaitGroup
	wg.Add(len(links))

	for i := range links {
		c[i] = make(chan string) // Initialize each channel
		go checkLink(links[i], c[i], &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	for _, ch := range c {
		fmt.Println(<-ch)
	}
}

func checkLink(link string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("All Not good with link :", link)
		c <- "Might be down I think"
		return
	}
	fmt.Println("All good with link :", link)
	c <- "All good with link"
}
