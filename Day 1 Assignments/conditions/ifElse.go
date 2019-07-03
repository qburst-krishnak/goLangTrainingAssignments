package main 

import . "fmt"

func main() {
	if 5 == 5.0 {
		Println("5 == 5.0")	
	}

	if 200%2 == 0 {
		Println("200 is even")	
	} else {
		Println("200 is odd")
	}

	if 5 < 0 {
		Println("5 is negative")	
	} else if 5%2 == 0 {
		Println("5 is even")	
	} else {
		Println("5 is odd")
	}
}
