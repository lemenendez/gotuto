package main

import (
	"fmt"
	"math"
)

type Heap struct {
	data []int
	len  int
}

func NewHeap() *Heap {
	h := &Heap{}
	h.len = 0
	return h
}

func (h *Heap) getParentIdx(i int) int {
	return (i - 1) / 2
}

func (h *Heap) getLeftChildIdx(i int) int {
	return i*2 + 1
}

func (h *Heap) getRightChildIdx(i int) int {
	return i*2 + 2
}
func (h *Heap) Add(val int) {
	h.data = append(h.data, val)
	h.len += 1
	i := h.len - 1
	for i != 0 {
		fidx := h.getParentIdx(i)
		if h.data[fidx] > h.data[i] {
			h.swap(i, fidx)
			i = fidx
		} else {
			break
		}
	}
}

func (h *Heap) Min() int {
	if h.len > 0 {
		return h.data[0]
	}
	return math.MaxInt64
}

/* http://www.mathcs.emory.edu/~cheung/Courses/171/Syllabus/9-BinTree/heap-delete.html
1, Delete a node from the array
       (this creates a "hole" and the tree is no longer "complete")

    2. Replace the deletion node
       with the "fartest right node" on the lowest level
       of the Binary Tree
       (This step makes the tree into a "complete binary tree")

    3. Heapify (fix the heap):

         if ( value in replacement node < its parent node )
            Filter the replacement node UP the binary tree
         else
	    Filter the replacement node DOWN the binary tree
*/
func (h *Heap) Delete(val int) {
	if h.len > 0 {
		// find val
		needle := -1
		for i := 0; i < h.len; i++ {
			if h.data[i] == val {
				needle = i
			}
		}
		if needle >= 0 {
			// replace it the last one (fartest right node)
			h.data[needle] = h.data[h.len-1]
			h.len -= h.len - 1
			// heapyfy
			pIdx := h.getParentIdx(needle)
			if h.data[needle] < h.data[h.getParentIdx(needle)] {
				// heapify up
				i := needle
				for i != 0 {
					h.swap(i, pIdx)
					i = pIdx
				}
			} else {
				// heapify down
			}

		}
	}
}

func (h *Heap) filterDown() {

}

func (h *Heap) swap(x, y int) {
	temp := h.data[y]
	h.data[y] = h.data[x]
	h.data[x] = temp
}

func main() {

	h := NewHeap()
	/*
				 1
			   /   \
			3	     6
		  /	  \    /
		5	   9  8


		**/
	h.Add(1)
	h.Add(3)
	h.Add(6)
	h.Add(5)
	h.Add(9)
	h.Add(8)

	fmt.Println(h)

	h1 := NewHeap()
	h1.Add(10)
	h1.Add(20)
	h1.Add(30)
	h1.Add(100)

	fmt.Println(h1)
}
