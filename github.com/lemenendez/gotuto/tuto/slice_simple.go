package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {

	myslice:=make([]int, 0, 3)
	
	var control string
	
	for control!="X" {
		fmt.Printf("Enter a integer number or X to exit\n")
		fmt.Scan(&control)

		newint, err:= strconv.Atoi(control)

		if err!= nil {
			continue 	// return earlier but not finish exec
		}	
		
		myslice = append(myslice, newint)
		sort.Ints(myslice)

		fmt.Printf("%v\n",myslice)
	}

}
