package main

import T "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}

	T.Println("Total :", total)
}

func main() {

	sum(1, 2, 3, 4)
	nums := []int{11, 12, 13, 14}
	sum(nums...)
}
