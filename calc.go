package main

import "fmt"

func main() {

	fmt.Println(sum(2, 2))
	fmt.Println(sub(10, 6))
}

func sum(a int, b int) int {

	return a + b
}

func sub(a int, b int) int {

	return a - b
}

func mult(a int, b int) int {

	return a * b
}

func div(a int, b int) int {

	return a / b
}
