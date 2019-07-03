package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])
	fmt.Println("Length:", len(s))

	s = append(s, "D")
	s = append(s, "E", "F")
	fmt.Println("After Appending :", s)
}
