package main

import . "fmt"

func main() {
	var i int = 1
	for i < 3 {
		Println(i)
		i++	
	}

	for j:=3; j<7; j++ {
		Println(j)	
	}

	for {
		Println("Infinite loop!!")
		break
	}

	for k:=1; k<6; k++ {
		if k%2 == 0 {
			continue
		}
		Println(k)
	}
}
