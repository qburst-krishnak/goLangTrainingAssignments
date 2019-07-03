package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	init1 := initSeq()

	fmt.Println(init1())
	fmt.Println(init1())
	fmt.Println(init1())

	init2 := initSeq()

	fmt.Println(init2())
	fmt.Println(init2())
	fmt.Println(init2())
}
