/*
Outplace: using another structure
Given an array a of n items, suppose we want an array that holds the same elements in reversed order and to dispose of the original. One seemingly simple way to do this is to create a new array of equal size, fill it with copies from a in the appropriate order and then delete a. 
*/

package main

import ("fmt")

func reverse(a []int) {
	length := len(a)
	half := (len(a) - 1) / 2
	for i := 0; i<half; i = i + 1 {
		tmp := a[i]
		a[i] =  a[length-1-i]
		a[length-1-i] = tmp
	}
}

func main() {
	a := []int{1, 4, 8, 12, 16}
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
}