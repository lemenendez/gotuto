/*
Write a program which reads information from a file and represents it in a slice of structs. 
Assume that there is a text file which contains a series of names. 
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, 
fname for the first name, and lname for the last name. 
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. 
Your program will successively read each line of the text file and create a struct
which contains the first and last names found in the file. Each struct created will be added to a slice, 
and after all lines have been read from the file, 
your program will have a slice containing one struct for each line in the file. 
After reading all lines from the file, your program should iterate through your slice 
of structs and print the first and last names found in each struct.

Submit your source code for the program, “read.go”.

*/

package main


import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	fname string
	lname string
}



func main() {
	
	var filename string

	var Persons []Person

	fmt.Printf("Please enter the file name\n")
	fmt.Scan(&filename)

	fData, fErr:= ioutil.ReadFile(filename)
	if fErr==nil 	{		
		lines := strings.Split(string(fData),"\n")
		for i := 0; i < len(lines); i++ {			
			content := strings.Split(lines[i]," ")
			if len(content)>1 {
				var person Person				
				person.fname = content[0]
				person.lname = content[1]
				Persons = append(Persons, person)
			}			
		}
		fmt.Printf("Slice data: %v\n", Persons)
	}	else {
		
		panic(fErr)
	}
}
