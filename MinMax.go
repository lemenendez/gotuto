package main

import "fmt"

func min(arr []int, length int) int {
	minIdx := 0
	for i := 0; i < length; i++ {
		if arr[minIdx] > arr[i] {
			minIdx = i
		}
	}
	return arr[minIdx]
}

func max(arr []int, length int) int {
	minIdx := 0
	for i := 0; i < length; i++ {
		if arr[minIdx] < arr[i] {
			minIdx = i
		}
	}
	return arr[minIdx]
}

func minmax(a[]int, count int) (int, int) {
	minIdx:=0
	maxIdex:=0
	for i := 0; i < count; i++ {
		if a[minIdx] > a[i] {
			minIdx = i
		}else if(a[maxIdex]<a[i]) {
			maxIdex = i
		}	
	}	
	return a[minIdx], a[maxIdex]
}

func main() {
	scores := []int{60, 50, 95, 80, 70}
	length := len(scores)
	min := min(scores, length)
	max := max(scores, length)
	fmt.Printf("Min Value=%d\n", min)
	fmt.Printf("Max Value=%d\n", max)
	min, max = minmax(scores, length)
	fmt.Printf("Min Value=%d\n", min)
	fmt.Printf("Max Value=%d\n", max)
}
