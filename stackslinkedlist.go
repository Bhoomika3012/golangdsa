package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stacklist struct {
	top *Node
}

func (s *stacklist) push(n int) {
	newNode := &Node{data: n}
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *stacklist) pop() {
	if s.top == nil {
		fmt.Println("stack is empty")
	} else {
		temp := s.top
		s.top = temp.next
		temp = nil
	}
}
func (s *stacklist) stacktop() {
	if s.top == nil {
		fmt.Println("stack is empty")
	} else {
		fmt.Println(s.top.data)
	}
}

func (s *stacklist) display() {
	p := s.top
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

func main() {
	l := stacklist{}
	l.push(10)
	l.push(20)
	l.push(30)
	//l.display()
	//l.pop()
	l.stacktop()
	l.display()
}
