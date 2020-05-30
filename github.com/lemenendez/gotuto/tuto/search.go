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
