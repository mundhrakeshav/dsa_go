package stack

import (
	"testing"
)

// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T) {
	var stackArr []int = createNew(5, 1, 2, 43, 5)
	if len(stackArr) != cap(stackArr) {
		t.Error("Len Cap Mismatch")
	}
	t.Run("Push", func(t *testing.T) {
		initialLen := len(stackArr);
		stackArr = stackPush(stackArr, 23);
		if len(stackArr) <= initialLen {
			t.Error("Did not push")
		}

	})
	t.Run("Pop", func(t *testing.T) {
		initialLen := len(stackArr);
		// var stackArr []int;
		pop, stackArr := stackPop(stackArr)

		if len(stackArr) >= initialLen {
			t.Error("Did not pop")
		}

		if pop != 23 {
			t.Error("Incorrect Pop")
		}
	})
}
