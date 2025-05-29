package main

import "fmt"

func init() {
    fmt.Println("Inside init function")
}
func init() {
    fmt.Println("Inside init function2")
}
func init() {
    fmt.Println("Inside init function3")
}

func main() {
    fmt.Println("Inside main function")
}

// ✅ You don’t call init() yourself. Go runs it automatically.

// Use Case	Description
// 🔐 Initializing config/env	Load from env or files
// 🗃️ Setting up database drivers	Before app starts
// 🧪 Test preparation	Setup state for tests
// 🔄 Registering dependencies	Like plugin registration or metrics