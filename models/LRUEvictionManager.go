package models

type LRUEvictionManager struct {
	head, tail *Node
}

func (l *LRUEvictionManager) push(key string) {
	node := &Node{value: key}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *LRUEvictionManager) pop() string {
	deletedData := l.head.value
	l.head = l.head.next
	l.head.prev = nil
	return deletedData
}

func (l *LRUEvictionManager) clear() {
	// var i int = 0
	// for l.head != nil {
	// 	i++
	// 	fmt.Println(l.head)
	// 	l.head = l.head.next
	// 	fmt.Println(i)
	// }
	l.head = nil
	l.tail = nil
}
