package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite loop!")
		break // Exit after one iteration
	}

	numbers := []int{1, 2, 3, 4, 5}

	for i, num := range numbers {
		fmt.Println(i, num)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end // Jump to the labeled statement
		}
		fmt.Println(i)
	}

end:
	fmt.Println("Loop ended")

	array := make([][]int, 5)
	for i := range array {
		array[i] = make([]int, 5)
	}

	// Initialize the array with zeros
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		array[i][j] = 0
	//	}
	//}

	// Print the array
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}

	dayOfWeek := time.Now().Weekday()
	switch dayOfWeek {
	case time.Saturday:
		fmt.Println("It's Saturday!")
	case time.Sunday:
		fmt.Println("It's Sunday!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Type switch
	var value interface{}
	value = 10
	switch value.(type) {
	case int:
		fmt.Println("value is an integer")
	case string:
		fmt.Println("value is a string")
	default:
		fmt.Println("value is of unknown type")
	}

	// Multiple values
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Excellent!")
	case 'B', 'C':
		fmt.Println("Good job!")
	case 'D':
		fmt.Println("You can do better.")
	default:
		fmt.Println("Invalid grade.")
	}

	fmt.Println(factorial(5))

	fmt.Println(fibonacci(5))

	memoizedFactorial := memoize(factorial)
	memoizedFibonacci := memoize(fibonacci)

	// Calculate and print factorials
	for i := 0; i <= 5; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, memoizedFactorial(i))
	}

	// Calculate and print Fibonacci numbers
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci number %d: %d\n", i, memoizedFibonacci(i))
	}

	//if ok retrun
	m := map[string]int{"key1": 10, "key2": 20}
	value, ok := getValue(m, "key1")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}

	// Example 2: Using convertToInt
	var num interface{} = 10.5
	intNum, ok := convertToInt(num)
	if ok {
		fmt.Println("Converted to int:", intNum)
	} else {
		fmt.Println("Conversion failed")
	}

	// Example 3: Function Check
	var s MyStruct

	// Check if s implements HasProcess
	if hasProcessMethod1(s) {
		fmt.Println("s implements HasProcess") // This will be printed
	} else {
		fmt.Println("s doesn't implement HasProcess")
	}

	if hasProcessMethod2(s) {
		fmt.Println("s implements 2") // This will be printed
	} else {
		fmt.Println("s doesn't implement 2")
	}

	if hasProcessMethod3(s) {
		fmt.Println("s implements 3")
	} else {
		fmt.Println("s doesn't implement 3") // This will be printed
	}
}

// half baked
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			fmt.Println("Got Value from cache !! ", v)
			return v
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// if ok pattern
func getValue(m map[string]int, key string) (int, bool) {
	value, ok := m[key]
	if !ok {
		return 0, false // Return default value and false if key not found
	}
	return value, true
}

func convertToInt(value interface{}) (int, bool) {
	num, ok := value.(int)
	if !ok {
		return 0, false // Return 0 and false if conversion fails
	}
	return num, true
}

type HasProcess interface {
	process()
}

// Define a struct that implements HasProcess
type MyStruct struct{}

// define associate funtion
func (MyStruct) process() {
	fmt.Println("funtion to check")
}

// Function to check if a variable implements HasProcess
func hasProcessMethod1(i interface{}) bool {
	_, ok := i.(HasProcess)
	return ok
}

func hasProcessMethod2(i interface{}) bool {
	_, ok := i.(interface{ process() })
	return ok
}

func hasProcessMethod3(i interface{}) bool {
	_, ok := i.(interface{ process2() })
	return ok
}
