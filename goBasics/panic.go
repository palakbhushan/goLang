//  In Go, panic is a built-in function that immediately stops the normal execution of a program. 
// It’s how Go handles unexpected, unrecoverable errors.

package main
import "fmt"

func main() {

	// panic(interface{})

	// Example of a valid input
	process(10)

	// Example of an invalid input
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After Panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}