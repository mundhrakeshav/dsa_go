package queue

import (
	"errors"
	"fmt"
)


type Node[T any] struct {
	val T;
	next *Node[T];
}

type QueueLinkedList[T any] struct {
	len int;
	head *Node[T];
	tail *Node[T];
}

func createQueueLinkedList[T any](_initalData T) QueueLinkedList[T] {
	
	newNode := &Node[T]{
		val: _initalData,
	}	
	return QueueLinkedList[T]{
		len: 1,
		head: newNode,
		tail: newNode,
	}
}


func enqueueLinkedList[T any](queue *QueueLinkedList[T], val T) QueueLinkedList[T] {
	
	if queue.head == nil {
		return createQueueLinkedList(val)
	}
	
	
	newNode := &Node[T]{
		val: val,
	}



	queue.tail.next = newNode 
	queue.tail = newNode;
	queue.len++;
	return *queue;
}


func dequeueLinkedList[T any](queue *QueueLinkedList[T]) (QueueLinkedList[T], error)  {

	if queue.head == nil {
		fmt.Println("RAN 1")
		return *queue, errors.New("Underflow");
	}
	
	if queue.head == queue.tail {
		fmt.Println("RAN 2")
		queue.head = nil;
		queue.tail = nil;
		return *queue, nil;
	}

	queue.head = queue.head.next;
	queue.len--;
	return *queue, nil
}

func printQueue[T any](queue QueueLinkedList[T])  {
	node := queue.head;
	if node == nil {
		return
	}
	for i := 0; i < queue.len; i++ {
		fmt.Println(node.val)
		node = node.next
	}
}