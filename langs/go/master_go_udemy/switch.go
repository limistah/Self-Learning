package main

import (
	"fmt"
	"time"
)

func main () {

	language := "golang"
	switch language {
	case "Python":
		 fmt.Println("You are learning Pythoin! You don't use curly braces but indentation")
	case "Go", "golang":
		 fmt.Println("Good, go for Go! You are using curly braces {}!")
	}

	// Can be used to determine the value of an expression
	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!!")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

}