package tree

const (
	height = 5
)

type TreeNode struct {
	Keywords [height - 1]int //键值，用来划分区间数据
	Children []*TreeNode     // 保存叶子结点指针
}

type TreeLeafNode struct {
	k int
}
