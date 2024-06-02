package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWrite struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
	//fmt.Printf("%+v", resp)

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(logWrite{}, resp.Body)

}

func (logWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote every thing ", len(bs))
	return len(bs), nil
}
