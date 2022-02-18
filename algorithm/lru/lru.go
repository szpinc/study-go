package lru

type Node struct {
	Key, Val   int
	Next, Prev *Node
}

type DoubleList struct {
	Head, Tail *Node // 头，尾结点
	Size       int   // 链表大小
}

// AddLast 在链表末尾添加节点，时间复杂度: O(1)
func (list *DoubleList) AddLast(node *Node) {
	node.Prev = list.Tail.Prev
	node.Next = list.Tail

	list.Tail.Next = node
	list.Tail = node
	list.Size++
}

// Remove 移除指定节点
func (list *DoubleList) Remove(node *Node) {

}

// RemoveFirst 移除第一个
func (list *DoubleList) RemoveFirst() {

}
