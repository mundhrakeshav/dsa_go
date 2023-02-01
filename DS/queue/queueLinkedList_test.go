package queue

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestLinkedListQueueArray(t *testing.T) {
	queueLinkedList := createQueueLinkedList(5)

	t.Run("Push", func(t *testing.T) {
		printQueue(queueLinkedList)
		fmt.Println("++++++++")
		enqueueLinkedList(&queueLinkedList, 12)
		enqueueLinkedList(&queueLinkedList, 15)
		enqueueLinkedList(&queueLinkedList, 134)
		enqueueLinkedList(&queueLinkedList, 1432)
		enqueueLinkedList(&queueLinkedList, 165322)
		printQueue(queueLinkedList)
		fmt.Println("++++++++")
		dequeueLinkedList(&queueLinkedList);
		dequeueLinkedList(&queueLinkedList);
		dequeueLinkedList(&queueLinkedList);
		dequeueLinkedList(&queueLinkedList);
		dequeueLinkedList(&queueLinkedList);
		dequeueLinkedList(&queueLinkedList);
		printQueue(queueLinkedList)
	})

}
