package main

import "fmt"

func main() {

	a ,b := 4,5

	x, y := 3.14, 20.4

	var q,w float32 = 3.11, 22.22

	swapInt(&a, &b)
	swapFloat(&x, &y)
	swapFloat32(&q, &w)

	fmt.Printf("%v% v\n", a, b)								// print only
	fmt.Printf("%v %v\n", x, y)								// print with format
	fmt.Printf("%v %v\n", q, w)								// print with format
	

}
