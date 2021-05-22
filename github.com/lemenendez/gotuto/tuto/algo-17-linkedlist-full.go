package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l *LinkedList) Len() int {
	return l.length
}

func (l LinkedList) Display() {
	i := l.head
	for i != nil {
		fmt.Printf("%v ", i.data)
		i = i.next
	}
	fmt.Println()
}

func (l *LinkedList) Add(data int) {
	node := &Node{data, nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *LinkedList) DelAt(position int) {
	if l == nil {
		return
	}
	if position == 0 && l.head != nil {
		l.head = l.head.next
		return
	}
	iter := l.head
	var prev *Node
	i := 0
	for iter != nil {
		if i == position {
			prev.next = iter.next
			return
		}
		prev = iter
		iter = iter.next
		i += 1
	}
}

func (l *LinkedList) InsertAt(position int, data int) {
	if position < 0 || position > l.length-1 {
		return
	}
	if l.head == nil && position != 0 {
		return
	}
	if l.head == nil {
		node := &Node{data, nil}
		l.head = node
		l.tail = node
	} else {
		counter := 0
		iter := l.head
		var prev *Node
		for iter != nil {
			if counter == position {
				node := &Node{data, nil}
				if prev == nil {
					node.next = l.head
					l.head = node
					if node.next != nil {
						l.tail = node.next
					} else {
						l.tail = node
					}
				} else {
					prev.next = node
					return
				}
			}
			counter += 1
			prev = iter
			iter = iter.next
		}
	}
}

func main() {
	l := &LinkedList{}
	l.Add(3)
	l.Add(4)
	l.Add(20)
	// l.Display()
	l.DelAt(2)
	// l.Display()
	l.DelAt(0)
	l.DelAt(0)
	l.DelAt(0)
	l.Display()
	l.InsertAt(0, 1)
	// l.Display()
	l.InsertAt(0, 4)
	l.Display()
}
