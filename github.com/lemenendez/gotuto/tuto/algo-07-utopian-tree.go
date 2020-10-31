/*
The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height. Each summer, its height increases by 1 meter.
A Utopian Tree sapling with a height of 1 meter is planted at the onset of spring. How tall will the tree be aftergrowth cycles?

For example, if the number of growth cycles is, the calculations are as follows:

Period  Height
0          1
1          2
2          3
3          6
4          7
5          14


*/

package main

import (
    "fmt"
)


func utopianTree(n int32) int32 {
	h:=0
	len := int(n)
	for i:=0 ; i<=len; i++ {
		if i % 2 == 0 {
			h++
		} else {
			h = h *2
		}
		fmt.Printf("%v\t %v\n", i, h)
	}
	return int32(h)
}

func main() {
	utopianTree(5)
}
