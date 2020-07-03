package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func genSort(origin []int, out chan []int) {
	l := len(origin)
	output := make([]int, l)
	copy(output, origin)
	sort.Ints(output)
	out <- output
}

func split(actions []int) [][]int {
	batchSize := len(actions) / 4
	if batchSize<4 {
		batchSize = 3
	}
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	return batches
}

func sortedInsert(a []int, v int) []int {
	l := len(a)
	if l == 0 {
		return append(a, v)
	}

	spot := 0
	found := false
	for i := 0; i < l; i++ {
		if a[i] > v {
			spot = i
			found = true
			break
		}
	}
	if found {
		tmp := append([]int{}, a[0:spot]...)
		tmp = append(tmp, v)
		return append(tmp, a[spot:]...)
	} else {
		return append(a, v)
	}

}

func GetInputs() []int {
	var nums []int = make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a list of integer numbers separated by space:\n")
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	for i := 0; i < len(tokens); i++ {
		if s, err := strconv.Atoi(tokens[i]); err == nil {
			nums = append(nums, s)
		}
	}
	return nums
}

func main() {

	// 1: handling inputs
	input := GetInputs()

	// 2: end result slice
	output := make([]int, 0)

	// 3: setting communication channel
	ch := make(chan []int)
	calls := 0
	
	// 4: split the work into batches
	actions := split(input)

	// 5: calling go routines, for each batch start up a go routine
	
	for i := 0; i < len(actions); i++ {		
		// fmt.Printf("go routine will sort:%v\n", actions[i])
		go genSort(actions[i], ch)
		calls += 1
	}

	// 6: merging results
	var tmp []int
	for i := 1; i <= calls; i++ {
		tmp = <-ch
		// fmt.Printf("incoming array:%v\n", tmp)
		for i := 0; i < len(tmp); i++ {
			// maitain a sorted array
			output = sortedInsert(output, tmp[i])
			// fmt.Printf("input: %v sorted:%v\n", tmp[i], output)
		}
	}

	fmt.Printf("sorted:%v\n", output)
}
