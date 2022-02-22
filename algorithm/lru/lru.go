package lru

type (
	Node struct {
		Key, Val   int
		Next, Prev *Node
	}
	DoubleList struct {
		Head, Tail *Node // 头，尾结点
		Size       int   // 链表大小
	}
	Cache struct {
		Map   map[int]*Node
		Cache *DoubleList
		Cap   int
	}
)

func newDoubleList() *DoubleList {
	head := Node{
		Key: 0,
		Val: 0,
	}
	tail := Node{
		Key: 0,
		Val: 0,
	}

	head.Next = &tail
	tail.Prev = &head

	return &DoubleList{
		Head: &head,
		Tail: &tail,
		Size: 0,
	}
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
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
}

// RemoveFirst 移除第一个
func (list *DoubleList) RemoveFirst() *Node {
	if list.Size == 0 {
		return nil
	}
	list.Remove(list.Head)
	return list.Head
}

func NewCache(cap int) *Cache {
	m := make(map[int]*Node)

	return &Cache{
		Map:   m,
		Cache: newDoubleList(),
	}
}

func makeRecently(cache *Cache, key int) {
	node := cache.Map[key]
	if node == nil {
		return
	}
	// 删除节点
	cache.Cache.Remove(node)
	// 在末尾增加节点
	cache.Cache.AddLast(node)
}

func addRecently(cache *Cache, key int, val int) {
	node := Node{
		Key: key,
		Val: val,
	}
	cache.Cache.AddLast(&node)
	cache.Map[key] = &node
}
