package stack

type Node[T any] struct {
	Val T;
	Next *Node[T];
}

type LinkedListStack[T any] struct {
	Head *Node[T];
	Length int;
}

func CreateNewLinkedList[T any](val T) *LinkedListStack[T] {
	head := Node[T]{
		Val: val,
	};

	return &LinkedListStack[T]{
		Head: &head,
		Length: 1,
	}
}

func (ll *LinkedListStack[T]) Push(Val T)  {
	newHead := &Node[T]{
		Val: Val,
		Next: ll.Head,
	}

	ll.Head = newHead;
	ll.Length++;
}

func (ll *LinkedListStack[T]) Print()  {
	node := ll.Head;
	for i := 0; i < ll.Length; i++ {
		node = node.Next;
	}

}

func (ll *LinkedListStack[T]) Pop()  {
	if ll.Head == nil{
		return;
	}
	
	if ll.Head.Next == nil {
		ll.Head = nil;
	} else {
		ll.Head = ll.Head.Next;
	}
	ll.Length--;
}