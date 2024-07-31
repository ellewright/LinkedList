package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T] // tail
	value T
}

type LinkedList[T any] struct {
	head   *Node[T]
	length uint
}

func (ll *LinkedList[T]) index(idx uint) (T, bool) {
	// return value and "ok", indicating if index is valid
	var value T

	if idx >= ll.length {
		return value, false
	}

	currentIdx := 0
	current := ll.head
	for current != ll.head || currentIdx == 0 {
		if idx >= uint(currentIdx) {
			value = current.value
			break
		}

		currentIdx++
		current = current.next
	}

	return value, true
}

func (ll *LinkedList[T]) append(value T) {
	ll.length++
	if ll.length == 1 {
		ll.head = &Node[T]{}
		ll.head.prev = ll.head
		ll.head.next = ll.head
		ll.head.value = value
		return
	}

	node := &Node[T]{}
	node.value = value

	tail := ll.head.prev
	tail.next = node
	node.prev = tail
	node.next = ll.head
	ll.head.prev = node
}

func (ll *LinkedList[T]) pop() T {
	ll.length--
	if ll.length == 0 {
		var null T
		value := ll.head.value
		ll.head.value = null
		return value
	}

	tail := ll.head.prev
	tailPrev := tail.prev
	ll.head.prev = tailPrev
	tailPrev.next = ll.head

	return tail.value
}

func (ll *LinkedList[T]) printList() {
	values := []T{}

	current := ll.head
	for ll.length != 0 {
		values = append(values, current.value)
		current = current.next
		if current == ll.head {
			break
		}

		fmt.Println(values)
	}
}

func main() {
	ll := LinkedList[int]{}
	ll.append(10)
	ll.append(15)
	ll.append(4)
	ll.append(9)
	ll.append(4)
	ll.append(9)

	ll.printList()
	fmt.Println(ll.index(7))
}
