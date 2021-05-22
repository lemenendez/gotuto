/* Permutations and combinations definitions */

package main

import (
	"fmt"
	"errors"
)

func factorial(n int) (uint64, error) {
	if n<0 {
		return 0, errors.New("input is negative")
	}
	var ret uint64 = 1
	for i:=1; i<=n; i++ {
		ret *= uint64(i)
	}
	return ret, nil
}

func main() {

	input:="CAT"

	fmt.Printf("Given a n amount of elements, there are n! ways of sort them: ordereds or permutations:\nCAT\nCTA\nATC\nACT\nTAC\nTCA\n")
	length:=len(input)
	res, err := factorial(length)
	if err==nil {
		fmt.Printf("For input: %s, there are %v\n", input, res)
	}else {
		fmt.Printf("there was an error %v", err)
	}
}