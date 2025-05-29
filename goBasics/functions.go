package main
import (
	"fmt"
)

func main() {

	//just a naming conventions
	//private fun starts with lower case
	// public func name start with uppercase -- like Println

	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	// sum := add(1, 2)
	// fmt.Println(add(2, 3))

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }

	// greet()

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))


	//returning multiple return values
	// func functionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2,...){
		//code block
		// return returvalue1, returnValue2,...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}


}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}