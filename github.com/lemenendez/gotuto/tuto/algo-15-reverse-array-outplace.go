/*
Outplace: using another structure
Given an array a of n items, suppose we want an array that holds the same elements in reversed order and to dispose of the original. One seemingly simple way to do this is to create a new array of equal size, fill it with copies from a in the appropriate order and then delete a. 
*/

package main

import ("fmt")

func reverse(a []int) []int {
	r:=make([]int, len(a))
	for i, j := 0, len(a)-1; j>=0; i, j = i+1, j-1 {
		// fmt.Printf("%v, %v\n", i, j)
		r[i] = a[j]
	}
	return r
}

func main() {
	a := []int{1, 4, 8, 12, 16}
	fmt.Print(reverse(a))
}