package main

import "fmt"

type node struct {
	data int
	next *node
}

type singllylist struct {
	head *node
	tail *node
}

func (s *singllylist) push(n int) {
	newnode := node{data: n}
	newnode.next = s.head
	s.head = &newnode
}
func (s *singllylist) pushany(n, Pos int) {
	newnode := node{data: n}
	p := s.head
	for i := 1; i < Pos-1; i++ {
		p = p.next
	}
	newnode.next = p.next
	p.next = &newnode
}

func (s *singllylist) pushlast(n int) {
	newnode := node{data: n}
	p := s.head
	for p.next != nil {
		p = p.next
	}
	p.next = &newnode
}

func (s *singllylist) removefirst() {
	p := s.head
	s.head = p.next
	p.next = nil
}
func (s *singllylist) removeany(pos int) {
	p := s.head.next
	prev := s.head
	for i := 1; i < pos-1; i++ {
		p = p.next
		prev = prev.next
	}
	prev.next = p.next
}
func (s *singllylist) removelast() {
	p := s.head.next
	prev := s.head

	for p.next != nil {
		p = p.next
		prev = prev.next
	}
	prev.next = nil
}
func (s *singllylist) display() {
	if s.head == nil {
		fmt.Println("linked list is empty")
	} else {
		p := s.head
		for p != nil {
			fmt.Println(p.data, "-->")
			p = p.next
		}

	}

}
func (s *singllylist) len() int {
	p := s.head
	count := 0
	for p != nil {
		count += 1
		p = p.next
	}
	return count

}
func main() {
	l := singllylist{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.pushany(25, 2)
	l.pushlast(50)
	//l.removefirst()
	l.removeany(2)
	l.removelast()
	l.display()
	fmt.Println(l.len())

}
