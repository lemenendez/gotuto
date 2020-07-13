/*

Write a Bubble Sort program in Go. 
The program should prompt the user to type in a sequence 
of up to 10 integers. 
The program should print the integers out on one line, 
in sorted order, from least to greatest. 
Use your favorite search tool to find a description 
of how the bubble sort algorithm works.

As part of this program, you should write a 
function called BubbleSort() which takes a slice 
of integers as an argument and returns nothing. 
The BubbleSort() function should modify the slice 
so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is 
the Swap operation which swaps the position of two 
adjacent elements in the slice. 
You should write a Swap() function which performs 
this operation. Your Swap() function should 
take two arguments, a slice of integers and an 
index value i which indicates a position in the slice. 
The Swap() function should return nothing, 
but it should swap the contents of the slice in 
position i with the contents in position i+1.

*/

package main

import (
	"fmt"
)

func BubbleSort(origin []int) {
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
			}
		}
	}	
}

func main() {
	var max int = 10
	var origin []int = make([]int,0, max)
	var num int
	for i := 0; i < max; i++ {
		fmt.Printf("Enter Number %v\n", i+1)
		fmt.Scan(&num)
		origin = append(origin,num)
	}
	BubbleSort(origin)
	fmt.Printf("Sorted Numbers:%v\n", origin)


}
