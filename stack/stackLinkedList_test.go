package stack

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestLinkedListArray(t *testing.T) {
	stackLinkedList := createNewLinkedList(5)

	t.Run("Push", func(t *testing.T) {
		// pushToLinkedListStack(stackLinkedList, 3)
		// pushToLinkedListStack(stackLinkedList, 1)
		// pushToLinkedListStack(stackLinkedList, 34)
		fmt.Println(stackLinkedList.head)
		// printLinkedListStack(stackLinkedList)
		popLinkedListStack(stackLinkedList)
		popLinkedListStack(stackLinkedList)
		fmt.Println("=========")
		printLinkedListStack(stackLinkedList)
	})

}
