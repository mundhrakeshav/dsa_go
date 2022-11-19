package linkedlist

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestLinkedList(t *testing.T) {

	cyclic := createNew(5)
	t.Run("Add, Delete and Visualise", func(t *testing.T) {
		// initialLen := len(stackArr);
		cyclic.addToLinkedList(12)
		cyclic.addToLinkedList(11)
		cyclic.addToLinkedList(111)
		cyclic.addToLinkedList(65)
		cyclic.addToLinkedList(34)
		cyclic.printLinkedList()
		fmt.Println("++++++++++++")
		cyclic.deleteHead()
		cyclic.deleteHead()
		cyclic.printLinkedList()
		fmt.Println("++++++++++++")
		cyclic.deleteHead()
		cyclic.deleteHead()
		cyclic.deleteHead()
		cyclic.printLinkedList()
		cyclic.deleteHead()
		cyclic.deleteHead()
	})
}
