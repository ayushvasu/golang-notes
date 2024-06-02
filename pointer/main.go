package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	fmt.Println(namePointer)
	fmt.Println(*namePointer)
	fmt.Println(*&namePointer) // same as "namePointer"
	printPointer(namePointer)
}

func printPointer(namePointer *string) { // because copy is created for address thats why different value than line 10
	fmt.Println(&namePointer)
}
