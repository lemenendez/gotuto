package main

import ("fmt")


func fibonnaci(v int) int {
	if v<=0 {
		return 0
	}
	if v ==1 {
		return 1
	}
	return fibonnaci(v-1) + fibonnaci(v-2)
}

func main() {
	fmt.Println(fibonnaci(0))
	fmt.Println(fibonnaci(1))
	fmt.Println(fibonnaci(2))
	fmt.Println(fibonnaci(3))
	fmt.Println(fibonnaci(4))
	fmt.Println(fibonnaci(5))
	fmt.Println(fibonnaci(6))
	fmt.Println(fibonnaci(7))
	fmt.Println(fibonnaci(8))
	fmt.Println(fibonnaci(9))
	fmt.Println(fibonnaci(10))
	fmt.Println(fibonnaci(11))
	fmt.Println(fibonnaci(12))

}