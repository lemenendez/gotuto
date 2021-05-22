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

func (l *LinkedList) PushBack(n *Node) {
	if n == nil {
		return
	}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.length++
}

func (l *LinkedList) Delete(position int) {
	if l == nil {
		return
	}
	if position == 0 {
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

func main() {
	l := &LinkedList{}
	l.PushBack(&Node{3, nil})
	l.PushBack(&Node{4, nil})
	l.PushBack(&Node{20, nil})
	l.Display()
	l.Delete(2)
	l.Display()
}
