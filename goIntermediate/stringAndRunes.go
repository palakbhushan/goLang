// | Use Case                    | Why Use Runes                            |
// | --------------------------- | ---------------------------------------- |
// | Unicode-safe string parsing | Multilingual support (हिन्दी, 中文, emoji) |
// | Tokenizers in compilers     | Handle code points, not bytes            |
// | Text UI apps (CLI, TUI)     | Align multi-byte characters correctly    |
// | Regex/token scanning        | Need char-by-char decoding               |


// | Term     | Type    | Description                         |
// | -------- | ------- | ----------------------------------- |
// | `byte`   | uint8   | One raw byte                        |
// | `rune`   | int32   | One Unicode code point (UTF-8 safe) |
// | `string` | \[]byte | UTF-8 encoded sequence of bytes     |

// A rune in Go is just an alias for int32 — it represents a Unicode code point.
// Why?
// Because Go strings are UTF-8 encoded, which means:

// 1 character might be 1 to 4 bytes.

// You need runes to handle multibyte characters correctly.


package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, Go!" // Go!lo
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
