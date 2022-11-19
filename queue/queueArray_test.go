package queue

import (
	"testing"
)

func TestQueueArray(t *testing.T) {
	var queueArr []int = createNew(5, 1, 2, 43, 5)
	if len(queueArr) != cap(queueArr) {
		t.Error("Len Cap Mismatch")
	}
	t.Run("EnQueue", func(t *testing.T) {
		initialLen := len(queueArr);
		queueArr = enqueue(queueArr, 34);
		if len(queueArr) <= initialLen {
			t.Error("Did not Enqueue")
		}
		if len(queueArr) != 6 {
			t.Error("Enqueue error")
		}
	})

	t.Run("DeQueue", func(t *testing.T) {
		initialLen := len(queueArr);
		queueArr = dequeue(queueArr);
		if len(queueArr) >= initialLen {
			t.Error("Did not Enqueue")
		}
		if len(queueArr) != 5 {
			t.Error("Enqueue error")
		}
	})

}
