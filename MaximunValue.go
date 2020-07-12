package main

import "fmt"

func max(arrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] {
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}
	max:=arrays[length-1]
	return max
}

func main() {

	var scores = []int{60, 50, 95, 80, 70}

	var length = len(scores)

	var max = max(scores, length)

	fmt.Printf("Max value = %d\n", max)

}