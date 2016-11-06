package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// NOTE: I'd much rather just create a new Scanner and have a Split func read each int, but it breaks
	// Hackerrank's input method (and works locally) so I'll just do this ugly stuff to make it work.
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	num_arr := make([]int, n)

	for i, v := range arr {
		num, _ := strconv.Atoi(v)
		num_arr[i] = num
	}

	swaps, first, last := sort(num_arr)
	fmt.Print("Array is sorted in ", swaps)
	fmt.Print(" swaps.\n")
	fmt.Println("First Element:", first)
	fmt.Println("Last Element:", last)
}
