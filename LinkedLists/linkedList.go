package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Print() {
	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Printf("Current node: %d \n", current.data)
		current = current.next
	}
}

func (ll *LinkedList) Remove(index int) {
	if index == 0 {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for i := 1; i < index-1; i++ {
		current = current.next
	}
	current.next = current.next.next
}

func (ll *LinkedList) Get(index int) int {
	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return -1
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data
}

func (ll *LinkedList) Set(index int, data int) {
	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.data = data
}

func (ll *LinkedList) Length() int {
	if ll.head == nil {
		return 0
	}
	current := ll.head
	length := 1
	for current.next != nil {
		current = current.next
		length++
	}
	return length
}

func (ll *LinkedList) Clear() {
	ll.head = nil
}

func main() {
	var ll LinkedList
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Print()
	fmt.Println("Length: ", ll.Length())
	ll.Remove(0)
	ll.Remove(1)
	ll.Remove(2)
	ll.Remove(3)
	ll.Remove(4)
	ll.Remove(5)
	ll.Print()
	fmt.Println("Length: ", ll.Length())
	ll.Clear()
	ll.Print()
}
