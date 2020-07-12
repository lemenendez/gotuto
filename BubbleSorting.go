package main

import "fmt"

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
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}