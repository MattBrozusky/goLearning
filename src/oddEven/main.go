package main

import "fmt"

func main() {
	numSet := []int{}

	for i := 0; i <= 10; i++ {
		numSet = append(numSet, i)
	}

	for _, num := range numSet {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
