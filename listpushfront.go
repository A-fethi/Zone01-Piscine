package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}
	if l.Head == nil {
		newNode := &NodeL{Data: data, Next: nil}
		l.Head = newNode
		return
	}
	l.Head = newNode
}
