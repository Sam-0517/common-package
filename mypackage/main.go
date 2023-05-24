package mypackage

import "fmt"

func Add(a, b int) int {
	c := a + b
	return c
}

func Name(name string) {
	fmt.Printf("Name is %v", name)
}
