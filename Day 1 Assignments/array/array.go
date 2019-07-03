package main

import . "fmt"

func main() {
	var x [5]int
	Println("X :", x)

	x[1] = 01
	x[4] = 04
	Println("X :", x)
	Println("Length of X :", len(x))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	Println("2d: ", twoD)
}
