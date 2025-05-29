//  defer schedules a function call to run after the surrounding function returns.
// func main() {
//     fmt.Println("Start")
//     defer fmt.Println("Deferred")
//     fmt.Println("End")
// }
// Output:
// Start
// End
// Deferred

// How does defer behave with goroutines?
// 🧠 Key Rule:
// defer applies to the current function, not to the goroutine.
// So if you use defer inside a goroutine, it runs when that goroutine finishes — not when the main function finishes.

// func task() {
//     defer fmt.Println("Cleanup in goroutine")
//     fmt.Println("Working in goroutine")
// }

// func main() {
//     go task()  // go routine 
//     time.Sleep(1 * time.Second)
//     fmt.Println("Main done")
// }

// defer is triggered at the end of task(), which is inside the goroutine.





// ❌ defer in main won't clean up goroutines
// func task() {
//     fmt.Println("Goroutine started")
// }

// func main() {
//     defer fmt.Println("Main exiting")
//     go task()
// }
// Wrong assumption: defer waits for goroutine — ❌ It doesn't.

// Why?

// The main function can finish before the goroutine even starts.

// And defer in main() will run when main() ends — it doesn't wait for go task().

package main
import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Deffered i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}