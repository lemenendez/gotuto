package main

import (
	"fmt"
)


func main() {
	var text string

	fmt.Print("Enter text:")
	fmt.Scan(&text)

	if Search(text) {
		fmt.Printf("Found\n")
	}	else {
		fmt.Printf("Not Found!\n")
	}

}
