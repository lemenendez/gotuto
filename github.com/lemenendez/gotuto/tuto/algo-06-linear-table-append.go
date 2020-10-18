/*

From: Graphic Go Algorithms
Credits: Yan Hu
Name: Linear Table Append

Add a X score to the end of the array


*/

package main

import "fmt"

func print(a[] int, count int) {
	for i:=0; i< count; i++ {
		if i==count-1 {
			fmt.Printf("%d\n", a[i])
		} else {
			fmt.Printf("%d, ", a[i])
		}
	}
}

func main() {

	scores := []int{90, 70, 50, 80, 60, 85}
	count := len(scores)
	tmp := make([]int, count + 1)
	for i := 0; i < count; i++ {
		tmp[i] = scores[i]
	}
	tmp[count]= 75
	scores = tmp
	// there is an error in the book  Yan Hu
	print(scores, len(scores))
}
