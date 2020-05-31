package main

import (
	"fmt"
	"strconv"
	"sort"
)

// stores the an int value and if a flat to check if it has alread a value
type NillableInt struct {
	value int
	has bool
}

func NewNillableInt(value int, has bool) NillableInt {
	nillableInt := NillableInt{}
	nillableInt.value = value
	nillableInt.has = has
	return nillableInt
}

// not need but it might be useful
func NewNillable(value int) NillableInt {
	nillableInt := NillableInt{}
	nillableInt.value = value
	nillableInt.has = true
	return nillableInt
}

// serialize as string
func (v NillableInt) String() string {
	if !v.has {
		return ""
	}		
	return fmt.Sprintf("%v", v.value)
}

// it updates the value if doesnt have any value
func SetValue(s *[]NillableInt, v int) {
	foundSlot := false
	for i := 0; i < len((*s)); i++ {
		if !(*s)[i].has {
			(*s)[i].value = v
			(*s)[i].has = true
			foundSlot = true
			break
		}
	}
	if !foundSlot {
		*s = append((*s), NillableInt{v, true})
	}
}

func main() {

	// slice with a length of 3
	myslice:=make([]NillableInt, 3)
	
	var control string

	// main loop, exits only if X is entered by the user
	for control!="X" {
		fmt.Printf("Enter a integer number or X to exit\n")
		fmt.Scan(&control)

		newint, err:= strconv.Atoi(control)

		if err!= nil {
			continue 	// return earlier but not finish exec
		}	
		
		SetValue(&myslice, newint)
		// sort
		sort.Slice(myslice, func(i, j int) bool {
			return myslice[i].value > myslice[j].value
		})

		fmt.Printf("%v\n",myslice)
	}

}
