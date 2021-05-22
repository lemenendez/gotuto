package main

import (
	"fmt"
)

// brute force
func m(arr []int) bool {
	val_transition := 0
	for i := 1; i < len(arr)-1; i++ {
		/* detect the following cases:
		// 	#1 	\ _ _
		//  #2    _ _
		//  #4  /
		*/
		if arr[i] == arr[i-1] || arr[i] == arr[i+1] {
			return false
		}
		/* detect the following cases:
		//  #5  /     \		can happen only once
		*/
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			val_transition++
		}
		/* detect the following cases:
		// 	#9 	\ _ /			cannot happen
		*/
		if arr[i-1] > arr[i] && arr[i] < arr[i+1] {
			return false
		}
	}

	return val_transition == 1
}

func main() {

	fmt.Print(m([]int{0, 3, 2, 1}))
	fmt.Print(m([]int{2, 1}))
	fmt.Print(m([]int{3, 5, 5}))

	fmt.Print(m([]int{0, 1, 2, 1, 2}))
}
