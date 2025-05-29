// A closure is a function value that references variables from outside its body. 
// It can "close over" those variables — even after the outer function has finished executing.

// In simpler terms: a closure remembers the environment it was created in.

// 🔁 Think of Closures as:
// Functions + Memory
// They don’t forget the environment they were created in.

package main

import "fmt"

func main() {

	// sequence := adder()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtracter := func() func(int) int {

		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// Using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

// 4. Encapsulating Private State in Games or Bots
// func scoreTracker() func(int) int {
//     score := 0
//     return func(points int) int {
//         score += points
//         return score
//     }
// }

// func main() {
//     playerScore := scoreTracker()
//     fmt.Println(playerScore(10)) // 10
//     fmt.Println(playerScore(5))  // 15
// }