package dst

// MakeBSTree 从提供的分片中构建一个二叉排序树
func MakeBSTree(datas []BTSortDataer) *BTNode {
	var bst *BTNode
	for k := range datas {
		if k == 0 {
			bst = &BTNode{Data: datas[k]}
		} else {
			bst.BSTInsert(datas[k])
		}
	}
	return bst
}

// BSTSearch 二叉搜索树查找，类似先序遍历
// 如果找到 返回第一个值为其parent节点
func (bst *BTNode) BSTSearch(data BTSortDataer) (*BTNode, *BTNode) {
	var p = bst
	ret := bstsearch(data, bst, &p)
	return p, ret
}

func bstsearch(data BTSortDataer, bst *BTNode, p **BTNode) *BTNode {
	if bst != nil {
		switch bst.Data.Compare(data) {
		case 0:
			return bst
		case 1:
			*p = bst
			return bstsearch(data, bst.Lchild, p)
		case -1:
			*p = bst
			return bstsearch(data, bst.Rchild, p)
		}
	}
	return nil
}

// BSTInsert 插入一个节点，成功返回true,失败返回false
func (bst *BTNode) BSTInsert(data BTSortDataer) bool {
	if bst == nil {
		return false
	}
	parent, node := bst.BSTSearch(data)
	if node != nil {
		return false
	}
	if parent.Data.Compare(data) == 1 {
		parent.Lchild = &BTNode{Data: data}
	} else {
		parent.Rchild = &BTNode{Data: data}
	}
	return true
}

// BSTDelete 删除一个节点
func (bst *BTNode) BSTDelete(data BTSortDataer) *BTNode {
	if bst == nil {
		return nil
	}
	parent, node := bst.BSTSearch(data)
	// 未找到
	if node == nil {
		return nil
	}

	// 如果节点只有左孩子或者只有右孩子，子承父业
	if node.Lchild == nil {
		if parent.Lchild == node {
			parent.Lchild = node.Rchild
		} else {
			parent.Rchild = node.Rchild
		}
	} else if node.Rchild == nil {
		if parent.Lchild == node {
			parent.Lchild = node.Lchild
		} else {
			parent.Rchild = node.Lchild
		}
	} else { // 如果两个孩子都有，找到node的直接前驱
		parentPrevNode := node
		preNode := node.Lchild
		for preNode.Rchild != nil { // 此时prevNode只有左孩子，并且是node的前驱
			parentPrevNode = preNode
			preNode = preNode.Rchild
		}
		node.Data = preNode.Data    // 用preNode的值来代替node的值
		if parentPrevNode == node { // 如果前驱直接是node的左孩子，则将node的左孩子替换为前驱的左孩子
			node.Lchild = preNode.Lchild
		} else { // 如果前驱不是node的左孩子（隔着很多层）
			parentPrevNode.Rchild = preNode.Lchild
		}

		preNode.Lchild = nil
		return preNode
	}

	return node
}
