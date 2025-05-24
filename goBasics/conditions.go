package main
import "fmt"

func main(){
	age:=34

	if age >=1 {
		fmt.Println("passed")
	}

	score := 95

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
	
	if 10%2 == 0 || 6%2 == 0 {
		fmt.Println("Either 10 or 6 are even.")
	}

	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 and 6 are even.")
	}