/*

Starting with a 1-indexed array of zeros and a list of operations, for each operation add a value to each the array element between two given indices, inclusive. Once all operations have been performed, return the maximum value in the array.

Example

n = 10
queries = [ [1 5 3] [4 8 7] [6 9 1]]

Queries are interpreted as follows:
    a b k
    1 5 3
    4 8 7
    6 9 1



Add the values of k between the indices a and b inclusive:

index->	 1 2 3  4  5 6 7 8 9 10
		[0,0,0, 0, 0,0,0,0,0, 0]
		[3,3,3, 3, 3,0,0,0,0, 0]
		[3,3,3,10,10,7,7,7,0, 0]
		[3,3,3,10,10,8,8,8,1, 0]

The largest value is 10 after all operations are performed.

Function Description

Complete the function arrayManipulation in the editor below.

arrayManipulation has the following parameters:

    int n - the number of elements in the array
    int queries[q][3] - a two dimensional array of queries where each queries[i] contains three integers, a, b, and k.

Returns

    int - the maximum value in the resultant array

Input Format

The first line contains two space-separated integers n and m, the size of the array and the number of operations. Each of the next m lines contains three space-separated integers a,b and k, the left index, right index and summand.

Sample Input

5 3
1 2 100
2 5 100
3 4 100

Sample Output

200

Explanation

After the first update the list is 	100 100 0 	0 	0
After the second update list is 	100 200 100 100 100
After the third update list is 		100 200 200 200 100

The maximum value is 200.

*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func arrayManipulation(n int32, queries [][]int32) int64 {
	extremes := make(map[int32]int32)
	max := int64(-9223372036854775808)
	for i := 0; i < len(queries); i++ {
		low := queries[i][0]
		upper := queries[i][1]
		weight := queries[i][2]
		for j := low; j <= upper; j++ {
			extremes[j] = extremes[j] + weight
		}
		//extremes[low] = extremes[low] + weight
		// extremes[upper] = extremes[upper] + weight
	}
	for _, v := range extremes {
		max = getMax(max, int64(v))
	}

	return max
}

func getMax(val int64, newval int64) int64 {
	if int64(newval) > val {
		val = newval
	}
	return val
}

func main() {
	q := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}

	fmt.Println(arrayManipulation(5, q))
	q2 := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	fmt.Println(arrayManipulation(10, q2))

	q3 := [][]int32{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{3, 9, 15},
	}
	fmt.Println(arrayManipulation(10, q3))

	f, err := OpenFile("algo-21-input.txt")

	if err != nil {
		log.Fatal(err)
	}
	q4, err := ParseFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arrayManipulation(4000, q4))

}

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}

func ParseFile(f *os.File) ([][]int32, error) {

	reader := bufio.NewReader(f)

	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool
		var err error
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)
			if !isPrefix {
				break
			}
			if err != nil {
				if err != io.EOF {
					return nil, err
				}
				break
			}
		}
		line := buffer.String()
		if err == io.EOF {
			break
		}
	}

	return [][]int32{}, nil
}
