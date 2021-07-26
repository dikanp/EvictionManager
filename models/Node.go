package models

type Node struct {
	value string
	next, prev *Node
}

// func (n *Node) Next() *Node {
// 	return n.next
// }
// func (n *Node) Prev() *Node {
// 	return n.prev
// }

// func (n *Node) First() *Node {
// 	return n.h
// }