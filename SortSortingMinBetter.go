/*
Sort an array by repeatedly finding the minimum element
from unsorted part and putting it at the beggining.
This is "Better" or cleaner
*/

package main

import "fmt"

func min(a []int ) (int, int) {
	count:=len(a)
	minIdx, minVal := 0, a[0]	
	for i := 0; i < count-1; i++ {
		if minVal > a[i+1] {
			minVal = a[i+1]
			minIdx = i + 1
		}
	}
	return minIdx, minVal
}

func sort(a []int) {
	count := len(a)
	for i := 0; i < count; i++ {		
		minIdx, _:=min(a[i:])		// send "rest" part as slice
		minIdx = minIdx + i			// map index
		if i != minIdx {
			tmp := a[i]
			a[i] = a[minIdx]
			a[minIdx] = tmp
		}		
	}
}

func main() {

	scores := []int{90, 70, 50, 80, 60, 85}

	sort(scores)

	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d,", scores[i])
	}

}
