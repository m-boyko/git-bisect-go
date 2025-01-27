package main

import "fmt"

func Add(a, b int) int {
	return a
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(2, 2))
	fmt.Println(Sub(1, 2))
}
