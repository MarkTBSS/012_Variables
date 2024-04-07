package primitive

import "fmt"

func Primitive() {
	// Boolean
	var b bool = true
	fmt.Println("Boolean:", b)

	// String
	var s string = "Hello, Go!"
	fmt.Println("String:", s)

	// Integer
	var i int = 42
	fmt.Println("Integer:", i)

	// Unsigned Integer
	var ui uint = 42
	fmt.Println("Unsigned Integer:", ui)

	// Byte (Alias for uint8)
	var by byte = 'A'
	fmt.Println("Byte (uint8):", by)

	// Rune (Alias for int32)
	var r rune = 'ก'
	fmt.Println("Rune (int32):", r)

	// Float
	var f float32 = 3.14
	fmt.Println("Float:", f)

	// Complex
	var c complex64 = 1 + 2i
	fmt.Println("Complex:", c)
}
