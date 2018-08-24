package linklist

import (
	"testing"
)

func TestCycleCheck(t *testing.T) {
	head := &LinkNode{value: 1}
	tail := head
	for i := 2; i <= 20; i++ {
		tail.next = &LinkNode{value: i}
		tail = tail.next
	}

	if cycleIdx := CycleCheck(head); cycleIdx != 0 {
		t.Fatal(cycleIdx)
	}

	tail.next = head.next.next.next.next
	t.Log(tail.value, tail.next.value)
	if cycleIdx := CycleCheck(head); cycleIdx != 4 {
		t.Fatal(cycleIdx)
	}
}
