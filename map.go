package main

import (
	"fmt"
)


func main() {
	
	fmt.Println("mapping through int & strings")

	r := map[string]int {"Bob": 20, "Mary": 21} //map with string keys and int values (Key-Value pairs)

	fmt.Println("Bob = ", r["Bob"])
	fmt.Println("Mary = ", r["Mary"])
}
// This is a Go program that demonstrates how to use maps in Go.

// A map is a built-in data structure in Go that allows you to store key-value pairs. In this code snippet, a map named r is created with string keys and int values. The keys are "Bob" and "Mary", and the corresponding values are 20 and 21 respectively.

// The program then prints the values associated with the keys "Bob" and "Mary" using the map indexing syntax r["Bob"] and r["Mary"].