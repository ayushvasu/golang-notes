package main

import (
	"fmt"
	"strconv"
)

type contactInfo struct {
	email       string
	phoneNumber int64
}

type person struct {
	fName       string
	lName       string
	contactInfo // it could be with field name also  "contact contactInfo"
}

func (p person) toString() string {
	return p.fName + " " + p.lName + " can cantact by : " + p.contactInfo.email + ", " + strconv.FormatInt(p.phoneNumber, 10)
}

func (p *person) updateName(newName string) {
	(*p).fName = newName
}

func main() {
	ayush := person{"ayush", "vasu", contactInfo{"ayushvasu98@gmail.com", 8923375046}}

	fmt.Printf("%+v\n", ayush)
	(&ayush).updateName("vasu")

	fmt.Println(ayush.toString())

	ayush.updateName("ayush")

	fmt.Println(ayush.toString())
	//a := 1

	//fmt.Println(&a) //---> get the pointer out of variable
	//fmt.Println(*&a) //--> get value out of pointer
}
