package main

import "fmt"

func del(a, b int) int {
	return a - b
}

func main() {
	fmt.Printf("res is: %d", del(2, 1))
}