package main

import "fmt"

func main() {

	colors := map[string]string{
		"color": "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	colors2 := make(map[string]string)

	colors2["Ayush"] = "Vasu"

	colors2["Saxena"] = "Vasu"

	delete(colors2, "Saxena")

	fmt.Println(colors2)

	printMap(colors2)

}

func printMap(c map[string]string) {
	for a, b := range c {
		fmt.Println(a + " : " + b)
	}
}
