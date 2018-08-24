package linklist

// LinkNode 链表
type LinkNode struct {
	next  *LinkNode
	value int
}

// CycleCheck 是否有环 =0: 无， >0: 环起点位置
func CycleCheck(l1 *LinkNode) (cycleIdx int) {
	if l1 == nil {
		return
	}
	fast := l1
	slow := l1
	for {
		if fast.next != nil {
			fast = fast.next.next
		}
		if slow.next != nil {
			slow = slow.next
		}
		if fast == nil || slow == nil {
			return
		}
		if fast == slow {
			break
		}
	}
	slow = l1
	for {
		cycleIdx++
		slow = slow.next
		fast = fast.next
		if slow == fast {
			return
		}
	}
}
