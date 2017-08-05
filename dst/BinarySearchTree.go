package dst

import (
	"fmt"
)

// MakeBSTree 从提供的分片中构建一个二叉排序树
func MakeBSTree(datas []BTSortDataer) *BiTree {
	T := &BiTree{}
	for k := range datas {
		T.BSTInsert(datas[k])
	}
	return T
}

// BSTSearch 二叉搜索树查找，类似先序遍历
// 如果找到 返回节点和true
// 如果未找到，返回最后一次查找的节点，和false
func (T *BiTree) BSTSearch(data BTSortDataer) (*BTNode, bool) {
	var p *BTNode
	node := T.root
	for node != nil {
		p = node
		switch node.Data.Compare(data) {
		case 0:
			return node, true
		case 1:
			node = node.Lchild
		case -1:
			node = node.Rchild
		}
	}
	return p, false
}

// BSTInsert 插入一个节点，成功返回true,失败返回false(已存在)
func (T *BiTree) BSTInsert(data BTSortDataer) bool {
	parent, ok := T.BSTSearch(data)
	if ok {
		return false
	}

	// root
	if parent == nil {
		T.root = &BTNode{Data: data}
	} else {
		tmp := &BTNode{Data: data, Parent: parent}
		if parent.Data.Compare(data) == 1 {
			parent.Lchild = tmp
		} else {
			parent.Rchild = tmp
		}
	}
	return true
}

// BSTDelete 删除一个节点
func (T *BiTree) BSTDelete(data BTSortDataer) *BTNode {
	node, ok := T.BSTSearch(data)
	// 未找到
	if !ok {
		return nil
	}

	// 如果节点只有左孩子或者只有右孩子，子承父业
	if node.Lchild == nil || node.Rchild == nil {
		tmp := node.Lchild
		if tmp == nil {
			tmp = node.Rchild
		}
		if tmp != nil {
			tmp.Parent = node.Parent
		}
		// root
		if node.Parent == nil {
			T.root = tmp
		} else if node.Parent.Lchild == node {
			node.Parent.Lchild = tmp
		} else {
			node.Parent.Rchild = tmp
		}
	} else {
		// 如果两个孩子都有，找到node的直接后继
		successor := node.Rchild
		for successor.Lchild != nil {
			successor = successor.Lchild
		}
		if successor.Parent != node {
			T.TransPlant(successor, successor.Rchild)
			successor.Rchild = node.Rchild
			successor.Rchild.Parent = successor
		}
		T.TransPlant(node, successor)
		successor.Lchild = node.Lchild
		node.Lchild.Parent = successor
	}

	freeNode(node)
	return node
}

// CheckBSTree 检查一个树满不满足二叉搜索树
func CheckBSTree(T *BiTree) error {
	if err := CheckBiTree(T); err != nil {
		return err
	}

	if T.root == nil {
		return nil
	}

	var preNode *BTNode
	var err error

	T.NRMidOrder(func(node *BTNode) bool {
		if node.Lchild != nil {
			if node.Data.Compare(node.Lchild.Data) <= 0 {
				err = fmt.Errorf("node.left egt node, node: %v, node.left: %v", node.Data, node.Lchild.Data)
				return false
			}
		}
		if node.Rchild != nil {
			if node.Data.Compare(node.Rchild.Data) >= 0 {
				err = fmt.Errorf("node.right elt node, node: %v, node.right: %v", node.Data, node.Rchild.Data)
				return false
			}
		}
		if preNode != nil && preNode.Data.Compare(node.Data) >= 0 {
			err = fmt.Errorf("midOrder: preNode egt node, preNode: %v, node: %v", preNode.Data, node.Data)
			return false
		}
		preNode = node
		return true
	})

	return err
}
