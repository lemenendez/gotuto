package main

import "fmt"

func calcH(a []int32) int32 {
	h := int32(0)
	// l:=int32(len(a))
	for i := 0; i < len(a); i++ {
		h = h + a[i]
	}
	return h
}

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	he1 := calcH(h1)
	he2 := calcH(h2)
	he3 := calcH(h3)

	if he1 == he2 && he2 == he3 {
		return he1
	} else {
		for he1 != he2 && he2 != he3 || (len(h1) == 0 && len(h2) == 0 && len(h3) == 0) {
			fmt.Println(h1, h2, h3)
			fmt.Println(he1, he2, he3)
			// remove one from the heightes
			if he1 > he2 && he1 > he3 {
				h1 = h1[1:]
			} else if he2 > he1 && he2 > he3 {
				h2 = h2[1:]
			} else {
				h3 = h3[1:]
			}
			he1 = calcH(h1)
			he2 = calcH(h2)
			he3 = calcH(h3)
		}
		return he1
	}
}

func main() {

	fmt.Println(equalStacks([]int32{3, 2, 1, 1, 1}, []int32{4, 3, 2}, []int32{1, 1, 4, 1}))
}
