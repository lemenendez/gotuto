/*
From: Graphic Go Algorithms
Credits: Yan Hu
Name: Bubble Sorting Algorithm

Compare a[i] with a[i+1], if a[i] > a[i+1] are swaped
Remaining elements repeat this process, until sorting is completed

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

func Bubble(arr []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; i++ {
			if arr[j] > arr[j+1] {
				flag:=arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = flag
			}
		}
	}
}


func main() {
	scores:=[]int{90, 70, 50, 80, 60, 85}
	length:=len(scores)
	Bubble(scores, length)
	print(scores, length)
}