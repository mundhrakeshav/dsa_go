package linkedlist

import "fmt"

type Node[T any] struct {
	val T;
	next *Node[T];
	prev *Node[T];
}

type Cyclic[T any] struct {
	head *Node[T];
	len int;
}	

func createNew[T any] (initialVal T) (cyclic *Cyclic[T]) {
	newNode := &Node[T] {
		val: initialVal,
	}

	newNode.next = newNode;
	newNode.prev = newNode;

	cyclic = &Cyclic[T] {
		head: newNode,
		len: 1,
	}

	return cyclic;
}


func (cyclic *Cyclic[T]) addToLinkedList(val T)   {

	newNode := &Node[T]{
		val: val,
		next: cyclic.head,
		prev: cyclic.head.prev,
	}

	cyclic.head.prev.next = newNode;
	cyclic.head.prev = newNode;

	cyclic.len++;
}

func (cyclic *Cyclic[T]) deleteHead()   {

	if cyclic.head == nil {
		return;
	}

	if cyclic.head.next == cyclic.head {
		cyclic.head = nil;
		cyclic.len--;
		return;
	}

	cyclic.head.next.prev = cyclic.head.prev;
	cyclic.head.prev.next = cyclic.head.next;
	cyclic.head = cyclic.head.next;

	cyclic.len--;
}

func (cyclic *Cyclic[T]) printLinkedList()  {
	node := cyclic.head;
	for i := 0; i < cyclic.len; i++ {
		fmt.Println(node.val)
		node = node.next
	}
}