package main

import (
	"fmt"
)


func main(){

	a := 10
	b := &a
	fmt.Println(*b)
	fmt.Println(b)
	*b = 20
	fmt.Println(a)	
}

// This Go code snippet declares two variables a and b, where b stores the memory address of a. 
// Printing *b will give the value stored at that memory address, which is a. Changing the value through *b will reflect in a.