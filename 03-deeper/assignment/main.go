package main

import "fmt"

func main () {
	var oddEven []int

	for i := 0; i <= 10; i++ {
		oddEven = append(oddEven, i)
	}

	for _, val :=range oddEven {
		if val % 2 == 0 {
			fmt.Printf("%v is even", val)
			fmt.Println()
		} else
		{
			fmt.Printf("%v is odd", val)
			fmt.Println()
		}
		
	}
}