// 23,45,12,76,21
package main

import "fmt"

func twoMin(a []int) (int, int) {
	min1, min2 := 10000, 10000
	for i := 0; i < len(a); i++ {
		if a[i] < min1 {
			min1 = a[i]
		}
		if a[i] < min2 && a[i] > min1 {
			min2 = a[i]
		}
	}
	return min1, min2
}

func main() {
	a := []int{23, 45, 12, 76, 21}
	min1, min2 := twoMin(a)
	fmt.Printf("The two min for %v are %v and %v\n", a, min1, min2)

	b := []int{963, 330, 106, 212, 177, 489, 238, 900, 566, 261, 455, 415, 971, 513, 215, 917, 652, 717, 840, 325}                                                                                                                                                                                                     //  => 106, 177
	c := []int{9736, 9892, 9869, 9454, 3961, 2657, 6376, 2870, 4665, 8617, 7230, 4430, 8754, 2796, 6719, 6164, 1301, 5599, 601, 1737, 3569, 349, 2110, 6454, 8621, 1705, 9928, 1116, 551, 3486, 3193, 9002, 7082, 4777, 2225, 5852, 7486, 8422, 7370, 4549, 1162, 796, 4058, 7898, 5220, 9362, 7283, 1323, 6834, 2591} // => 349, 551

	min1, min2 = twoMin(b)

	fmt.Printf("The two min for %v are %v and %v\n", b, min1, min2)

	min1, min2 = twoMin(c)

	fmt.Printf("The two min for %v are %v and %v\n", c, min1, min2)

}
