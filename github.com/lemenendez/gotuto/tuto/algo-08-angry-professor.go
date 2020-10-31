/*
A Discrete Mathematics professor has a class of students. Frustrated with their lack of discipline, he decides to cancel class if fewer than some number of students are present when class starts. Arrival times go from on time (arrivalTime<=0) to arrived late (arrivalTime>0).

https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true

*/

package main

import (
    "fmt"
)


// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {
	ret :="NO"
	len:=len(a)
	ontime:=int32(0)
	for i := 0; i < len; i++ {
		if a[i] <= 0 {
			ontime++
		}
	}
	if ontime < k {
		ret = "YES"
	}
	return ret
}

func main() {
	students:=[]int32 {-1, -3, 4, 2}
	fmt.Printf("Cancelled?%v\n", angryProfessor(3, students))
}
