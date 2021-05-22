/*

https://en.wikipedia.org/wiki/Heap%27s_algorithm


*/

package main

import (
	"fmt"
)

func genPerm(int32 k, a[] int32, *[][]p) {
	
}

func main() {

	// A := [6]int32{5, 2, 4, 6, 1, 3}
	A := [2]int32{1, 2}
	perms := make([][]int32, 0)
	genPerm(A, &perms)
	fmt.Print(perms)


}