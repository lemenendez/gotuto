/*
Sort an array by repeatedly finding the minimum element
from unsorted part and putting it at the beggining.
*/

package main

import "fmt"

func sort(a []int) {
	count := len(a)
	for i := 0; i < count; i++ {		
		minIdx := i
		minVal := a[minIdx]
		// find minimum value of the "rest" of the array
		for j := i; j < count-1; j++ {
			if minVal > a[j+1] {
				minVal = a[j+1]
				minIdx = j + 1
			}
		}
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
		fmt.Printf("%d", scores[i])
	}

}
