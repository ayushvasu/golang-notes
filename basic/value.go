package main

import "os"

func getValue() string {
	return "Five of Diamonds"
}

func newCard() string {
	return "Five of Diamonds"
}

func saveToFile(str string) error {
	return os.WriteFile("FileName", []byte(str), 0644)
}
