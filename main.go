package main

import (
  "fmt"
)

func sort(array []int) []int {
  for {
    swaps := 0
    for i, v := range(array) {
      if i+1 == len(array) { continue }
      if v > array[i+1] {
        array[i], array[i+1] = array[i+1], array[i]
        swaps++
      }
    }
    if swaps == 0 { return array }
  }
}

func main() {
  arr := []int{2, 3, 4, 1, 1, 8, 3}
  fmt.Println(arr)
  sort(arr)
  fmt.Println(arr)
}
