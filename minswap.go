/*
You are given an unordered array consisting of consecutive integers [1, 2, 3, ..., n] without any duplicates. 
You are allowed to swap any two elements. You need to find the minimum number of swaps required to sort the array in ascending order. 

i   arr                     swap (indices)
0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
5   [1, 2, 3, 4, 5, 6, 7]

max distance between:
(0, 7) -> (1, 1)	(7-1)*(1-0)=6
(0, 7) -> (2, 3)	(7-3)*(2-0)=8
(0, 7) -> (3, 2)  (7-2)*(2-0)=10
(0, 7) -> (4, 5)  (7-4)*(4-0)=12
(0, 7) -> (5, 6)  0+7-5-6 = -4

*/

package main

import (
	"fmt"
)

func MinSwap(origin []int) {
	swaps:=0
	if (len(origin))<2 {
		return 
	}
	mustCheck := true
	for mustCheck {
		mustCheck = false
		for i := 0; i < len(origin)-1; i++ {
			var tmp int
			if origin[i]>origin[i+1] {
				tmp = origin[i]
				origin[i] = origin[i+1]
				origin[i+1] = tmp
				mustCheck = true
				swaps+=1
				fmt.Printf("swaps %v:%v\n", swaps, origin)
			}
		}
	}	
}

func main() {
	arr:= []int{7, 1, 3, 2, 4, 5, 6}

	fmt.Printf("Input Array:%v\n", arr)
	MinSwap(arr)



}
