/*
From: Graphic Go Algorithms
Credits: Yan Hu
Name: Linear Table Definition
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
	var scores = []int{90, 70, 50, 80 , 60, 85}

	var length = len(scores)

	print(scores, length)
}