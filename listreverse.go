package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	curr := l.Head
	var next *NodeL = nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Tail = l.Head
	l.Head = prev
}
