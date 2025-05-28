package main
import "fmt"

func main(){
	message := "Hello Universe"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, val: %c\n", i, v)
	}

	// range iterate maps in unordered way
	// range iterates array from index 0 to n
	// range iterate channels until it get closed
	
	for _, v := range message {  // if we are not using any value, this avoids memory leak
		// fmt.Println(i, v)
		fmt.Printf("val: %c\n", v)
	}
}