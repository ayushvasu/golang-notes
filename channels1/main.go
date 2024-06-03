package main

import (
	"fmt"
	"net/http"
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

	//c := make([]chan string, 5)//this only declare and initilize the array not initialize each channel
	//for i := range links {
	//	c[i] = make(chan string) // Initialize each channel
	//	go checkLink(links[i], c[i])
	//}

	c := make(chan string)
	//try to make this parallel
	for _, link := range links {
		go checkLink(link, c)
	}

	//for _, ch := range c {
	//	fmt.Println(<-ch)
	//}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("All Not good with link :", link)
		c <- "Might be down I think"
		return
	}
	fmt.Println("All good with link :", link)
	c <- "All good with link"
}
