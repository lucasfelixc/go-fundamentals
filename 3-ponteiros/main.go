package main

import "fmt"

func main() {
	a := 10
	b := &a // With &, b will look at the value of "a" in memory, "b" is a memory reference

	a = 100
	fmt.Println(*b) // With *, will show the value
}
