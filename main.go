package main

import (
	"fmt"
)

func sort(array []int) (int, int, int) {
	totalSwaps := 0
	firstElement := array[0]
	lastElement := array[len(array)-1]

	for {
		swaps := 0
		for i, v := range array {
			if i+1 == len(array) {
				lastElement = v
				continue
			}
			if v > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				if i == 0 {
					firstElement = array[0]
				}
				swaps++
			}
		}
		totalSwaps += swaps
		if swaps == 0 {
			return totalSwaps, firstElement, lastElement
		}
	}
}

func main() {
	arr := []int{2, 3, 4, 1, 1, 8, 3}
	fmt.Println(arr)
	swaps, first, last := sort(arr)
	fmt.Println(arr)
	fmt.Print("Array is sorted in ", swaps)
	fmt.Print(" swaps.\n")
	fmt.Println("First Element: ", first)
	fmt.Println("Last Element: ", last)
}
