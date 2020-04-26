package main

import "fmt"

func basicvardec1() {
	var myintvar int             // var declaration
	myintvar = 1                 // var assigment
	fmt.Printf("%v\n", myintvar) // printing
}

func basicvardec2() {

	myintvar := 1 // var declaration with type inference

	fmt.Printf("%v\n", myintvar)
}

func basicvardec3() {
	mybool := true

	fmt.Printf("%v\n", mybool)
}

func basicvardec4() {
	myfloat := 3.1 //float literal
	fmt.Printf("%v\n", myfloat)
}

func basicvardec5() {

	var i, j int = 1, 4 // multiple declaration

	fmt.Printf("%v\n%v\n", i, j)
}
