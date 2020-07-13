package main

import "fmt"

func print(a[] int) {
	var b int
	fmt.Printf("%v\n",a)
	fmt.Printf("Please press any key to continue\n")
	fmt.Scan(&b)
}

func sort(a []int, length int) {
	for i:=1; i < length; i++ {		// "better starts from position 1"
		insertElem := a[i]	// Take unsorted new elements
		insertPos := i		// Inserted position
		for j := insertPos - 1; j >=0; j-- {
			if insertElem < a[j] {
				a[j+1] = a[j]
				insertPos--
				print(a)
			}
		}
		a[insertPos] = insertElem
		print(a)
	}
}

func main() { 
	scores := []int{90, 70, 50, 80, 60, 85}
	length := len(scores)

	fmt.Printf("%v\n",scores)
	sort(scores, length)

	for i:=0; i<length; i++ {
		fmt.Printf("%d", scores[i])
	}

}