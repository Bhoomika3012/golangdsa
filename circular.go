package main

import "fmt"

type node struct {
	data int
	next *node
}

type circularlist struct {
	head *node
	tail *node
}

func (c *circularlist) push(n int) {
	newest := node{data: n}
	if c.head == nil {
		c.head = &newest
		c.tail = &newest
		newest.next = c.head
	} else {
		p := c.head
		newest.next = p

		c.head = &newest
		c.tail.next = c.head
	}
}
func (c *circularlist) pushany(n, pos int) {
	newest := node{data: n}
	p := c.head
	for i := 1; i < pos-1; i++ {
		p = p.next
	}
	newest.next = p.next
	p.next = &newest
}
func (c *circularlist) pushlast(n int) {
	newest := node{data: n}
	if c.head == nil {
		c.head = &newest
		c.tail = &newest
		newest.next = c.head
	} else {
		c.tail.next = &newest
		c.tail = &newest
		c.tail.next = c.head
	}
}
func (c *circularlist) removefirst() {
	if c.head != c.tail {
		c.head = c.head.next
		c.tail.next = c.head
	} else {
		if c.head == nil {
			c.tail = nil
		}
	}
}
func (c *circularlist) removeany(pos int) {
	p := c.head.next
	prev := c.head
	for i := 1; i < pos-1; i++ {
		p = p.next
		prev = prev.next
	}
	prev.next = p.next
}
func (c *circularlist) removelast() {
	p := c.head
	for p.next != c.tail {
		p = p.next
	}
	c.tail = p
	p = p.next
	c.tail.next = c.head
}
func (c *circularlist) display() {
	if c.head == nil {
		fmt.Println("linked list is empty")
	} else {
		p := c.head
		for p != nil {
			fmt.Printf("%v-->", p.data)
			p = p.next
			if p == c.head {
				break
			}
		}
	}
}
func (c *circularlist) len() int {
	p := c.head
	count := 0
	for p != nil {
		count += 1
		p = p.next
		if p == c.head {
			break
		}
	}
	return count
}
func main() {
	cl := circularlist{}
	cl.push(10)
	cl.push(20)
	cl.push(30)
	cl.pushlast(60)
	cl.pushany(40, 4)
	cl.pushany(50, 5)
	//cl.removefirst()
	//cl.removeany(2)
	//cl.removelast()
	cl.display()
	fmt.Println("\n", cl.len())
}
