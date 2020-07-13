/*
From: Graphic Go Algorithms
Credits: Yan Hu
Name: Maximum Value

Compare and swap elements in the array until the last
number in the array is the maximum

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

func max(a []int, length int) int {
	for i := 0; i < length-1; i++ {
		if a[i] > a[i+1] {
			var temp = a[i]
			a[i] = a[i+1]
			a[i+1] = temp
			print(a, length)
		}
	}
	max:=a[length-1]
	return max
}

func main() {

	var scores = []int{60, 50, 95, 80, 70}

	var length = len(scores)

	print(scores, length)

	var max = max(scores, length)

	fmt.Printf("Max value = %d\n", max)

}