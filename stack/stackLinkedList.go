package stack

import "fmt"

type Node[T any] struct {
	val T;
	next *Node[T];
}

type LinkedListStack[T any] struct {
	head *Node[T];
	length int;
}

func createNewLinkedList[T any](val T) *LinkedListStack[T] {
	head := Node[T]{
		val: val,
	};

	return &LinkedListStack[T]{
		head: &head,
		length: 1,
	}
}

func pushToLinkedListStack[T any](ll *LinkedListStack[T], val T)  {
	newHead := &Node[T]{
		val: val,
		next: ll.head,
	}

	ll.head = newHead;
	ll.length++;
}

func printLinkedListStack[T any](ll *LinkedListStack[T])  {
	node := ll.head;
	for i := 0; i < ll.length; i++ {
		fmt.Println(node.val)
		node = node.next;
	}

}

func popLinkedListStack[T any](ll *LinkedListStack[T])  {
	if ll.head == nil{
		return;
	}
	
	if ll.head.next == nil {
		ll.head = nil;
	} else {
		ll.head = ll.head.next;
	}
	ll.length--;
}