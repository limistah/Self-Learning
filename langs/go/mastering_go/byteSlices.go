// As Go does not have a char data type, it uses byte and rune for storing character values.
//A single byte can only store a single ASCII character whereas a rune can store Unicode characters.
//However, a rune can occupy multiple bytes.
package main

import "fmt"

func main() {
	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)
	// Print byte slice contents as text
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
	// Length of b
	fmt.Println("Length of b:", len(b))
}
