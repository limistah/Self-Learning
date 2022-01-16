package main

// File scope

// Package scope

// Local scope

import (
	"fmt"
	// import "fmt" gives an error, fmt is already imported
	f "fmt"
)

// is allowed, f becomes an alias


const done = false // package scoped

func main () {
	var task = "Running" // local (block) scoped, overshadows the package scope
	fmt.Println(task, done)
	f.Println("Bye Bye!")
}

func f1() {
	fmt.Printf("done in f1(): %v\n", done) // not overriden, would take the value of the package scope
}