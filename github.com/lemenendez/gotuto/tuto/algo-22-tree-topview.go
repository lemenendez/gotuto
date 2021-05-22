/*
Given a pointer to the root of a binary tree, print the top view of the binary tree.

The tree as seen from the top the nodes, is called the top view of the tree.

For example :

   1
    \
     2
      \
       5
      /  \
     3    6
      \
       4

Top View : 1->2->5->6
Complete the function topView and print the resulting values on a single line separated by space.

Constraints

 1 <= Nodes in the tree <= 500

Output Format

Print the values on a single line separated by space.

Sample Input

   1
    \
     2
      \
       5
      /  \
     3    6
      \
       4

Sample Output

1 2 5 6

Explanation

   1
    \
     2
      \
       5
      /  \
     3    6
      \
       4

From the top, only nodes 1, 2, 5, 6 are visible.


Maybe it needs to add detail about tree representation, i.e. that at any level left subtree never can overlap right subtree and vice versa. In the beginning I was confused - what is expected in following cases:

      1
     / \
    2   3
   / \
  4   5
       \
        6
         \
          7

Expected: 4-2-1-3-7 ? In fact - 4-2-1-3.

      1
     / \
    2   3
   / \
  4   5
     /
    6
   /
  7
 /
8

Expected: 8-4-2-1-3 ? In fact - 4-2-1-3, again.


*/

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type IntPtr []*int

func (n *Node) Insert(data int) *Node {
	if n == nil {
		return &Node{data, nil, nil}
	}
	if data <= n.data {
		n.left = n.left.Insert(data)
	} else {
		n.right = n.right.Insert(data)
	}
	return n
}

func (n *Node) InOrder() string {
	var r string
	if n != nil {
		r = n.left.InOrder() + " " + strconv.Itoa(n.data) + n.right.InOrder()
	}
	return r
}

func (n *Node) topViewLeft(a []int) []int {
	if n != nil {
		a = append(a, n.data)
		a = n.left.topViewLeft(a)
	}
	return a
}

func (n *Node) topViewRight(a []int) []int {
	if n != nil {
		a = append(a, n.data)
		a = n.right.topViewRight(a)
	}
	return a
}

func (n *Node) Print() {

}

func (n *Node) LevelOrder() []IntPtr {
	data := []IntPtr{}
	data = n.levelOrderRec(data, 0)
	return data
}

func (n *Node) levelOrderRec(data []IntPtr, level int) []IntPtr {
	if len(data) == level {
		data = append(data, IntPtr{})
	}
	if n != nil {
		// fmt.Println("adding:%v\n", n.data)
		data[level] = append(data[level], &n.data)
		data = n.left.levelOrderRec(data, level+1)
		data = n.right.levelOrderRec(data, level+1)
	} /*else {
		data[level] = append(data[level], nil)
	}*/
	return data
}

func (n *Node) TopView() []int {
	if n == nil {
		return []int{}
	} else {
		m := make(map[int][2]int)
		keys := []int{}
		res := []int{}
		n.topViewRec(m, 0, 0)

		for key := range m {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		for i := 0; i < len(keys); i++ {
			res = append(res, m[keys[i]][0])
		}
		return res
	}
}

func (n *Node) topViewRec(m map[int][2]int, column int, depth int) {
	if n != nil {
		dat := [2]int{n.data, depth}
		_, exists := m[column]
		if !exists {
			// add the very firs one we found
			m[column] = dat
			fmt.Printf("added %v in %v and %v\n", n.data, column, depth)
		} else {
			if m[column][1] > depth {
				m[column] = dat
				fmt.Printf("updated %v in %v and %v\n", n.data, column, depth)
			}
		}
		n.left.topViewRec(m, column-1, depth+1)
		n.right.topViewRec(m, column+1, depth+1)

	}
}

func (n *Node) topViewRightRec(m map[int]int, column int) {
	if n != nil {
		_, exists := m[column]
		if column <= 0 && !exists {
			// add the very firs one we found
			m[column] = n.data
			fmt.Printf("added %v in %v\n", n.data, column)
		} else if column > 0 {
			fmt.Printf("added/updated %v in %v\n", n.data, column)
			m[column] = n.data
		}
		n.right.topViewRightRec(m, column+1)
	}
}

func main() {
	var t1 *Node
	t1 = t1.Insert(1)
	t1 = t1.Insert(2)
	t1 = t1.Insert(3)

	out := t1.TopView()
	fmt.Println(out)

	// fmt.Println(t1.InOrder())
	var t2 *Node
	t2 = t2.Insert(2)
	t2 = t2.Insert(1)
	t2 = t2.Insert(3)
	out = t2.TopView()
	fmt.Println(out)
	// fmt.Println(t1.InOrder())

	var t3 *Node

	t3 = t3.Insert(1)
	t3 = t3.Insert(2)
	t3 = t3.Insert(5)
	t3 = t3.Insert(3)
	t3 = t3.Insert(6)
	t3 = t3.Insert(4)
	out = t3.TopView()
	fmt.Println(out)

	// fmt.Println(t2.InOrder())

	/// fmt.Println(t2.TopView())

	var t4 *Node

	dat, err := ioutil.ReadFile("/mnt/data/projects/gotuto/github.com/lemenendez/gotuto/tuto/algo-22-input.txt")
	if err == nil {
		data := string(string(dat[:]))
		nums := strings.Split(data, " ")
		for index := range nums {
			if num, err := strconv.Atoi(nums[index]); err == nil {
				// fmt.Printf("Inserting %v\n", num)
				t4 = t4.Insert(num)
			}
		}
	} else {
		fmt.Println(err)
	}
	// fmt.Println(t4.InOrder())
	// fmt.Println(t4.TopView())
	// fmt.Println(t4.InOrder())
	// out := t4.LevelOrder()
	out = t4.TopView()
	fmt.Println(out)
	//var left := []*int{}
	/*
		for i := 0; i < len(out); i++ {
			arr := out[i]
			//left := arr[0]
			//right := arr[len(arr)-1]
			fmt.Printf("[%d] of len %v:", i, len(arr))
			for j := 0; j < len(arr); j++ {
				if arr[j] != nil {
					fmt.Printf("%v ", *arr[j])
				} else {
					fmt.Printf("nil ")
				}
			}
			// fmt.Printf("%v ", i)
			/*if left != nil {
				fmt.Printf("%v ", *left)
			}
			if right != nil {
				fmt.Printf("%v ", *right)
			}
	*/
	//fmt.Println()

	/*for j := 0; j < len(arr); j++ {
		if arr[j] == nil {
			fmt.Print("nil ")
		} else {
			fmt.Printf("%04d ", *arr[j])
		}
	}*/
	//fmt.Println()
	//}
	// fmt.Println(t4.LevelOrder())
	/*
		f, err := OpenFile("algo-22-input.txt")
		if err != nil {
			ParseFile(f, t4)
		}
	*/
}

/*
func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}

func ParseFile(f *os.File, n *Node) error {

	reader := bufio.NewReader(f)
	// fmt.Print(reader)
	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool
		var err error
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)
			fmt.Print(1)
			if !isPrefix {
				break
			}
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
		}
		line := buffer.String()
		fmt.Print(line)
		nums := strings.Split(line, " ")
		for index := range nums {
			if num, err := strconv.Atoi(nums[index]); err == nil {
				fmt.Printf("Inserting %v\n", num)
				n.Insert(num)
			}
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}
*/
