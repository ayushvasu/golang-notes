package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fmt.Println(os.Args[1])

	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
