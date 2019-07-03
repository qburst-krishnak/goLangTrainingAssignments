package main

import T "fmt"

func returnMultipleValue() (int, int) {
	return 3, 1
}

func main() {
	x, y := returnMultipleValue()
	T.Println("X :", x)
	T.Println("Y :", y)

	_, v := returnMultipleValue()
	T.Println("V :", v)
}
