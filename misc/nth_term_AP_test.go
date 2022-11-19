package misc

import (
	"testing"
)

func TestQueueArray(t *testing.T) {
	t.Run("Nth Term AP", func(t *testing.T) {
		var a, n, d uint32 = 2, 5, 1
		z := nth_term_AP(a, n, d)
		if z != 6 {
			t.Error("Invalid Calculation of nth term")
		}
	})
}
