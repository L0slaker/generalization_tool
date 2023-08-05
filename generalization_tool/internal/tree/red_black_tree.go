package tree

import (
	"Prove/generalization_tool"
	"errors"
)

type color bool

const (
	Red   color = false
	Black color = true
)

var (
	ErrRBTreeSameRBNode     = errors.New("RBTree 不能添加重复节点Key")
	ErrRBTreeNotExistRBNode = errors.New("RBTree 不存在节点Key")
)

type rbNode[K any, V any] struct {
	key    K
	value  V
	color  color
	left   *rbNode[K, V]
	right  *rbNode[K, V]
	parent *rbNode[K, V]
}

func newRBNode[K any, V any](key K, value V) *rbNode[K, V] {
	return &rbNode[K, V]{
		key:    key,
		value:  value,
		color:  Red,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (n *rbNode[K, V]) setNode(v V) {
	if n == nil {
		return
	}
	n.value = v
}

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare generalization_tool.Comparator[K]
	size    int
}

func NewRBTree[K any, V any](compare generalization_tool.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		compare: compare,
		root:    nil,
	}
}

func (rb *RBTree[K, V]) Size() int {
	if rb == nil {
		return 0
	}
	return rb.size
}

// Add 增加节点
func (rb *RBTree[K, V]) Add(key K, value V) error {
	return rb.addNode(newRBNode(key, value))
}

func (rb *RBTree[K, V]) addNode(node *rbNode[K, V]) error {
	// 1.fixNode 存储要修正的节点
	var fixNode *rbNode[K, V]
	// 2.判断根节点是否为空，若为空则说明树为空，将新节点作为根节点，并赋值给fixNode
	if rb.root == nil {
		rb.root = newRBNode[K, V](node.key, node.value)
		fixNode = rb.root
	} else {
		// 3.如果根节点不为空，则从根节点开始遍历树，用变量t表示当前遍历到的节点
		// 变量cmp存储比较结果，变量parent存储当前节点的父节点
		t := rb.root
		cmp := 0
		parent := &rbNode[K, V]{}
		for t != nil {
			// 4.首先更新parent为当前节点t，然后比较新节点的键值和当前节点的键值大小
			// 并根据结果决定遍历左子树还是右子树，如果键值相同则报错
			parent = t
			cmp = rb.compare(node.key, t.key)
			if cmp < 0 {
				t = t.left
			} else if cmp > 0 {
				t = t.right
			} else {
				return ErrRBTreeSameRBNode
			}
		}
		// 5.当找到插入位置时，根据新节点的键值和父节点的关系构建新的节点 fixNode，
		// 将其赋值为红色
		fixNode = &rbNode[K, V]{
			key:    node.key,
			value:  node.value,
			color:  Red,
			parent: parent,
		}
		// 6.判断新节点是父节点的左子节点还是右子节点，并进行赋值操作
		if cmp < 0 {
			parent.left = fixNode
		} else {
			parent.right = fixNode
		}
	}
	// 7.增加红黑树的节点数量，然后调用 fixAfterAdd 修正红黑树
	rb.size++
	rb.fixAfterAdd(fixNode)
	return nil
}

// fixAfterAdd 插入时着色旋转
// 如果是空节点、root节点、父节点是黑无需构建
// 主要分为三种情况：
// 1.fixUncleRed 叔节点是红色右节点
// 2.fixAddLeftBlack 叔节点是黑色右节点
// 3.fixAddRightBlack 叔节点是黑色左节点
func (rb *RBTree[K, V]) fixAfterAdd(node *rbNode[K, V]) {

}

//// Delete 删除节点
//func (rb *RBTree[K, V]) Delete(key K) (V, bool) {
//
//}
//
//// Find 查找节点
//func (rb *RBTree[K, V]) Find(key K) (V, error) {
//
//}
