package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	// Escape inner double quotes
	fmt.Println("He says: \"Hello!\"")

	fmt.Println(`he says: "Hello!"`)

	s2 := "I like \n Go!"
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike
	`)

	fmt.Println(`C:\Users\Andrei`)
	fmt.Println(`C:\\Users\\Andrei`)

	// String Concatenation
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	// Go has two additional integer types called byte and rune that are aliases for unint8 and int32 data types. In Go, the byte and rune data types are used to distinguish characters from integer values.
	// Go doesn't have a char data type. It uses byte and rune to represent character values.

	var1, var2 := 'a', 'b'

	fmt.Printf("Type: %T, Value %d\n", var1, var2)

	str := "țară"

	// the lenght built in function displays the number of bytes not runes in a string
	// And a string is a slice of runes while each rune if of type int32
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // Logs the  bytes of each rune not the actual rune
	}
	// Decode a string rune by rune
	fmt.Println("\n" + strings.Repeat("#", 20))
	for i := 0; i < len(str); {
		var r, size = utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 20))
	// Automatically decode a string automatically
	for _, r := range str {
		fmt.Printf("%c", r)
	}

	// len returns the number of bytes in a string not the number of runes
	s1 = "Golang"
	fmt.Println(len(s1)) // ASCII letters occupy one byte returns 6

	fmt.Println(len("∂")) // unicode letter occupy more bytes returns 3

	// Get the number of runes/unicode points in a string
	n := utf8.RuneCountInString("∂") // 1
	fmt.Println(n)

	// Slicing a string
	// Slice returns a part of an string
	s1 = "I love Golang!"
	fmt.Println(s1[2:5]) // returns characters from 2 to 5

	s2 = "사랑해"           // I love you in Korean language
	fmt.Println(s2[0:2]) // Print the unicode value

	rs := []rune(s2)
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs[0]))

	// Converting between a string and a rune type is not efficient
	// A new backing array is created each time

	p := fmt.Println
	result := strings.Contains("I love Go Programming!", "lovex")
	p("Contains", result)

	result = strings.ContainsAny("success", "xys")
	p("ContainsAny", result)

	result = strings.ContainsRune("golang", 'g')
	p("Contains rune", result)

	n = strings.Count("cheese", "e")
	p("Count", n)

	p(strings.ToLower("Python, Golang java"))
	p(strings.ToUpper("Python, Golang java"))

	// comparing strings
	p("go" == "Go")                                   // When case does matter
	p(strings.ToLower("GO") == strings.ToLower("go")) // When case does not matter
	p(strings.EqualFold("Go", "go"))                  // a cleaner way for when case does not matter

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)
	myStr = strings.Replace("192.168.0.1", ".", ":", 1) // replace one
	p(myStr)
	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // replace all
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":") // replace all
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", myStr)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "xxxx")
	fmt.Printf("%#v\n", myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%v\n", fields)

	s1 = strings.TrimSpace("\t Goodbye Windows, Welcome Linux! \n")
	fmt.Printf("%q\n", s1)

	s2 = strings.Trim("...Hello Gophers!!!!?", ".!?")
	fmt.Printf("%q\n", s2)
}
