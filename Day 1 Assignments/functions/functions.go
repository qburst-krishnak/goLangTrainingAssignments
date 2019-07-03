package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func main() {
	a := add(10, 20)
	fmt.Println("10 + 20 :", a)

	b := sub(30, 10)
	fmt.Println("30 + 10 :", b)
}
