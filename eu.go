package main 

import "fmt"

/* 
func swap2(a int, b int) {
	tmp:= a
	a = b
	b=tmp
}
*/

/*
func divMod(numerator int, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return 
}
*/

func swap(a *int, b *int) {
	tmp:=*a
	*a = *b	
	*b = tmp
}

func Gcd(a,b int) int {
	if b>a {
		swap(&a, &b)
	}
	
	remainder := a % b
	// fmt.Println(remainder)
	if remainder==0 {
		return b
	}
		
	a = b
	b = remainder

	return Gcd(a,b)
}

func main() {

	fmt.Printf(" Gcd %v\n", Gcd(544, 119) )

}