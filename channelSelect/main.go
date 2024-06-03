package main

import (
	"fmt"
)

func main() {
	// Creating two channels
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Goroutine sending data to channel ch1
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Goroutine sending data to channel ch2
	go func() {
		for _, str := range []string{"hello", "world"} {
			ch2 <- str
		}
		close(ch2)
	}()

	// Goroutine receiving data from both channels
	go func() {
		for {
			select {
			case num, ok := <-ch1:
				if !ok {
					fmt.Println("Channel ch1 closed")
					//return
				}
				fmt.Println("Received from ch1:", num)
			case str, ok := <-ch2:
				if !ok {
					fmt.Println("Channel ch2 closed")
					//return
				}
				fmt.Println("Received from ch2:", str)
			}
		}
	}()

	// Wait for goroutines to finish
	var input string
	fmt.Println("Press enter to exit...")
	fmt.Scanln(&input)
}
