package main

import "fmt"

func binarySearch(a []int, v int) bool {
	low := 0
	up := len(a) - 1
	for low <= up {
		m := (low + up) / 2
		if a[m] < v {
			low = m + 1
		} else {
			up = m - 1
		}
	}
	if low == len(a) || a[low] != v {
		return false
	}
	return true
}

func main() {

	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	fmt.Print(binarySearch(items, 63))

}
