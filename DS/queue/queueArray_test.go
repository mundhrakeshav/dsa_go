package queue_test

import (
	queue "mundhrakeshav/dsa/DS/queue"
	"testing"
)

func TestCreateNew(t *testing.T) {
    q := queue.CreateNew(1, 2, 3)
    if len(q) != 3 {
        t.Errorf("Expected queue length to be 3, but got %d", len(q))
    }
}

func TestEnqueue(t *testing.T) {
    q := queue.CreateNew(1, 2, 3)
    q = q.Enqueue(4)
    if len(q) != 4 {
        t.Errorf("Expected queue length to be 4, but got %d", len(q))
    }
    if q[3] != 4 {
        t.Errorf("Expected last element to be 4, but got %v", q[3])
    }

    q = queue.CreateNew[int]()
    if len(q) != 0 {
        t.Errorf("Expected queue length to be 0, but got %d", len(q))
    }

}

func TestDequeue(t *testing.T) {
    q := queue.CreateNew(1, 2, 3)
    q = q.Dequeue()
    if len(q) != 2 {
        t.Errorf("Expected queue length to be 2, but got %d", len(q))
    }
    if q[0] != 2 {
        t.Errorf("Expected first element to be 2, but got %v", q[0])
    }
}