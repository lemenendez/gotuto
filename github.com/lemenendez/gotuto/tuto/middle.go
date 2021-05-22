/*
Given a random list of int values, find the maximum value, split in half and check if that number exists within the list if so return the value else find the closest number to the left or right in the list. Give priority to the right most value.
*/

// 40,33, 55,10,44,90
package main

import (
	"fmt"
	"math"
)

func middle(a []int) int {

	// find the max
	max := -1
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	half := max / 2
	dif := float64(10000)
	idx := -1
	for i := 0; i < len(a); i++ {
		if math.Abs(float64(half)-float64(a[i])) < dif {
			dif = math.Abs(float64(half) - float64(a[i]))
			idx = i
		}
		if a[i] == half {
			return half
		}
	}

	return a[idx]
}

func main() {

	a := []int{40, 33, 55, 10, 44, 90}

	b := []int{20, 40, 90, 10, 30} //  => 40

	c := []int{46, 44, 10, 40, 90, 20, 30} //  => 46

	d := []int{40, 46, 10, 20, 30, 44, 45, 90} // => 45

	e := []int{3535, 3636, 3296, 1496, 1183, 3692, 2420, 3495, 3133, 2510, 4773, 3501, 2442, 2129, 2949, 344, 3990, 1707, 1734, 530, 1788, 229, 4423, 2835, 4116, 2248, 1658, 4769, 925, 391} // => 2420

	fmt.Printf("arr: %v\n", a)
	fmt.Printf("max:%v\n", middle(a))

	fmt.Printf("Middle:for%v is %v\n", b, middle((b)))
	fmt.Printf("Middle:for%v is %v\n", c, middle((c)))
	fmt.Printf("Middle:for%v is %v\n", d, middle((d)))
	fmt.Printf("Middle:for%v is %v\n", e, middle((e)))

}
