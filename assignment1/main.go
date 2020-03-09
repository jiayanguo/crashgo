package main

import "fmt"

func main() {

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range data {
		if num%2 == 0 {
			fmt.Printf("%d is even \n", num)
		} else {
			fmt.Printf("%d is odd \n", num)
		}
	}
}
