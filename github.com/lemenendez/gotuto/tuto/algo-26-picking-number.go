/*
Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to .

*/
package main

import (
	"fmt"
	"math"
)

/*
 *
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// var set []int32
	max := int32(0)
	for i := 0; i < len(a); i++ {
		// start a new candidate set
		set := []int32{a[i]}
		// search a potencial element in the candidate set
		for j := 0; j < len(a); j++ {
			if i != j {
				fmt.Printf("set:%v, ele %v?", set, a[j])
				if compareAll(set, a[j]) {
					fmt.Printf("ok\n")
					set = append(set, a[j])
					if int32(len(set)) > max {
						max = int32(len(set))
					}
				} else {
					fmt.Printf("false\n")
				}
			}
		}
	}
	return max
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func compareAll(a []int32, needle int32) bool {
	for i := 0; i < len(a); i++ {
		if Abs(a[i]-needle) > 1 {
			return false
		}
	}
	return true
}

func minDiffAll(start int, end int, a []int32) bool {
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			if i != j {
				if math.Abs(float64(a[i]-a[j])) > 1 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	input := []int32{4, 6, 5, 3, 3, 1}
	fmt.Printf("picking numbers %v, should be 3\n", pickingNumbers(input))

	input = []int32{1, 2, 2, 3, 1, 2}
	fmt.Printf("picking numbers %v, should be 5\n", pickingNumbers(input))

}
