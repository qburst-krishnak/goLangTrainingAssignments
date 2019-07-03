package main

import . "fmt"
import . "time"

func main() {

	switch Now().Weekday() {
	case Friday:
		Println("Yo, its friday!")
	case Saturday, Sunday:
		Println("Its weekend")
	default:
		Println(":( Weekday :(")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
		    fmt.Println("I'm a bool")
		case int:
		    fmt.Println("I'm an int")
		case string:
		    fmt.Println("I'm a string")
		default:
		    fmt.Printf("Don't know type %T\n", t)
		}
    	}

	whatAmI(true)
	whatAmI(1)
    	whatAmI("Hello")
	WhatAmI(0.001)
}
