/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

Race conditions happen when there is two concurrent-parallel code that accesses the same resource (like a variable) at the same time.

In this example, routine 1 is increasing the value of p, and routine 2 decreases. Because go lang program cannot determinate 
when the context switch will happen there is no a way to determine the previous value of p without any sync mechanism

*/

package main

import (
	"fmt"
	"time"
)

func routine1(p *int) {
	for i:=0;i<8;i++ {
		*p++
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Routine 1: %v\n",*p)
	}
}

func routine2(p *int) {
	for i:=0;i<8;i++ {
		*p--
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Routine 2: %v\n",*p)
	}
}

func main() {
	myval:= 95

	go routine1(&myval)

	routine2(&myval)

}
