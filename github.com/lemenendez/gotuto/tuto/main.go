package main

import (
	"fmt"
)


func main() {

	var mynumber float32
	
	fmt.Printf("Enter a floating point number:\n");	

	fmt.Scan(&mynumber)

	fmt.Printf("%v\n", trunc(mynumber))

}
