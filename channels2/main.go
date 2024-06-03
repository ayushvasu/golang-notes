package main

import (
	"fmt"
	"net/http"
	"time"
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

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	//for {
	//	go checkLink(<-c, c)
	//}

	//same as avove
	for l := range c {
		//time.Sleep(5 * time.Second) //this is sleep on main go routine
		//go checkLink(l, c)

		//sleep with function literal

		go func(l string) { //to resolve the comment on 38 we added this and line 39
			time.Sleep(5 * time.Second)
			checkLink(l, c) //as l is out side variable from func literal it is showing the variable. and this might change before we use in literal
		}(l) //to resolve the comment on 38 we added this
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("All Not good with link :", link)
	} else {
		fmt.Println("All good with link :", link)
	}
	//time.Sleep(5 * time.Second)//with out function literal
	c <- link
}
