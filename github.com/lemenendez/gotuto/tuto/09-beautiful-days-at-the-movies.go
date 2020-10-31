/*
Lily likes to play games with integers. She has created a new game where she determines the difference between a number and its reverse. 

https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem



*/

package main

import (
	"fmt"
	"strconv"
	"math"
	
)
// so nice
// https://play.golang.org/p/AvJdDMwODA
// https://stackoverflow.com/questions/27287829/convert-byte-array-of-ascii-numeric-characters-to-int-using-go-language

func reverse(n int32 ) int32 {
	s := strconv.Itoa(int(n))
	len := len(s)
	res := 0
	for i := len-1; i >=0; i-- {		
		res *= 10
		res += int(s[i] - '0')
	}
	return int32(res)
}

func beautifulDays(i int32, j int32, k int32) int32 {
	b := int32(0)
	var div float64
	for i := i; i <= j; i++ {
		rev := reverse(i) 
		div = float64(i - rev) / float64(k) 
		// fmt.Printf("%v, %v %v\n", i, rev, div)
		if (div==math.Trunc(div)) {
			b++
		} 
	}
	return b
}


func main() {
	s:=reverse(800)
	fmt.Printf("%v\n",s)
	// 321

	fmt.Printf("Beautiful days %v\n", beautifulDays(13, 45, 3))

}