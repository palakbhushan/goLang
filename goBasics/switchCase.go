package main

import "fmt"

func main() {
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}