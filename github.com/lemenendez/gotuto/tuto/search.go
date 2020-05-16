package main

import (
	"strings"
)

func Search(text string) bool {
	lowered:= strings.TrimSpace(strings.ToLower(text))
	if strings.HasPrefix(lowered, "i") && strings.HasSuffix(lowered, "n") && strings.Contains(lowered, "a") {
		return true
	}	
	return false
}

/*
func main () {
	var text string
	var loweredtext string

	fmt.Print("Enter text:")
	fmt.Scan(&text)

	loweredtext = strings.TrimSpace(strings.ToLower(text))	

	if strings.HasPrefix(loweredtext, "i") && strings.HasSuffix(loweredtext, "n") && strings.Contains(loweredtext, "a") {
		fmt.Printf("Found!\n")
	}	else {
		fmt.Printf("Not Found!\n")
	}
}
*/