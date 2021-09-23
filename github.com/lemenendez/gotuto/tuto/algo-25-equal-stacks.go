package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
You have three stacks of cylinders where each cylinder has the same diameter, but they may vary in height. You can change the height of a stack by removing and discarding its topmost cylinder any number of times.

Find the maximum possible height of the stacks such that all of the stacks are exactly the same height. This means you must remove zero or more cylinders from the top of zero or more of the three stacks until they are all the same height, then return the height.

Example

h1 = [1, 2, 1, 1]
h2 = [1, 1, 2]
h3 = [1, 1]

There are 4,3 and 2 cylinders in the three stacks, with their heights in the three arrays. Remove the top 2 cylinders from h1 (heights = [1, 2]) and from h2 (heights = [1, 1]) so that the three stacks all are 2 units tall. Return 2 as the answer.

Note: An empty stack is still a stack.

Function Description

Complete the equalStacks function in the editor below.

equalStacks has the following parameters:

    int h1[n1]: the first array of heights
    int h2[n2]: the second array of heights
    int h3[n3]: the third array of heights

Returns

    int: the height of the stacks when they are equalized


*/

func calcH(a []int32) int32 {
	h := int32(0)
	for i := 0; i < len(a); i++ {
		h += a[i]
	}
	return h
}

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	done := false
	zero := int32(0)

	he1 := calcH(h1)
	he2 := calcH(h2)
	he3 := calcH(h3)

	for !done {
		fmt.Printf("h1:%v, h2:%v, h3:%v", he1, he2, he3)
		if he1 == he2 && he2 == he3 {
			return he1
		}
		if he1 == zero || he2 == zero || he3 == zero {
			return zero
		}
		// pop a cilinder from the heightest stack
		// and update the heigh
		if he1 >= he2 && he1 >= he3 {
			he1 -= h1[0]
			h1 = h1[1:]
			fmt.Printf(" h1>\n")
			continue
		} else if he2 >= he1 && he2 >= he3 {
			he2 -= h2[0]
			h2 = h2[1:]
			fmt.Printf(" h2>\n")
			continue
		} else {
			he3 -= h3[0]
			h3 = h3[1:]
			fmt.Printf(" h3>\n")
		}
	}
	return zero

}

func parseInts(line string) []int32 {
	var h []int32
	strInts := strings.Split(line, " ")
	for i := 0; i < len(strInts); i++ {
		item, err := strconv.ParseInt(strInts[i], 10, 64)
		if err != nil {
			panic(err)
		}
		h = append(h, int32(item))
	}
	return h
}

func main() {

	bytesRead, _ := ioutil.ReadFile("/mnt/data/projects/gotuto/github.com/lemenendez/gotuto/tuto/algo-25-input.txt")
	file_content := string(bytesRead)
	lines := strings.Split(file_content, "\n")
	h1 := parseInts(lines[1])
	h2 := parseInts(lines[2])
	h3 := parseInts(lines[3])

	fmt.Printf("heights h1:%v, h2:%v, h3:%v\n", calcH(h1), calcH(h2), calcH(h3))
	h := equalStacks(h1, h2, h3)

	fmt.Printf("height %v", h)
}
