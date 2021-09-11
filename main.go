package main

import "fmt"

var BuildID = "dev"

func main() {
	fmt.Printf("Build: %v\n", BuildID)
	fmt.Printf("2+2=%d\n", add(2, 2))
}

func add(a, b int) int {
	return a + b
}
