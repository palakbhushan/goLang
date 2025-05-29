//  What it does:
// Immediately terminates the program with the given exit code.

// Does not run any defer functions.

// Does not trigger panic or recover

package main

import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("This will NOT be printed")
    fmt.Println("Exiting...")
    os.Exit(1)
	fmt.Println("main func Exiting...")
}
// ✅ Defer is skipped, so you must not rely on cleanup in defer if using os.Exit.
// 🚨 panic() – Crash with Stack Trace
// 🔹 What it does:
// Stops execution by unwinding the stack.

// Runs all deferred functions in LIFO order.

// Can be recovered using recover() in a defer.