package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	current := l
	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next
	}
	if pos > count {
		return nil
	}
	return nil
}
