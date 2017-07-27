package dst

// BTSortDataer 可排序的数据定义
type BTSortDataer interface {
	Equal(data BTSortDataer) bool
	Compare(data BTSortDataer) int
}

// BTDStr string alias
type BTDStr string

// Equal implement BTSortDataer
func (btd BTDStr) Equal(b BTSortDataer) bool {
	return btd == b
}

// Compare implement BTSortDataer
func (btd BTDStr) Compare(b BTSortDataer) int {
	if btd == b {
		return 0
	}
	if btd > b.(BTDStr) {
		return 1
	}
	return -1
}

// BTDInt int alias
type BTDInt int

// Equal implement BTSortDataer
func (btd BTDInt) Equal(b BTSortDataer) bool {
	return btd == b
}

// Compare implement BTSortDataer
func (btd BTDInt) Compare(b BTSortDataer) int {
	if btd == b {
		return 0
	}
	if btd > b.(BTDInt) {
		return 1
	}
	return -1
}

// BTNode 二叉树节点
type BTNode struct {
	Data   BTSortDataer
	Lchild *BTNode
	Rchild *BTNode
}

type nrTmpNode struct {
	node     *BTNode
	visitNum int // 0:第一次访问，1：第二次访问
}

// PreSearch 先序查找指定值是否在树中，
// 如果找到 返回值所对应的节点
func (T *BTNode) PreSearch(data BTSortDataer) (*BTNode, bool) {
	if T == nil {
		return nil, false
	}
	if T.Data.Equal(data) {
		return T, true
	}
	n, found := T.Lchild.PreSearch(data)
	if found {
		return n, found
	}
	n, found = T.Rchild.PreSearch(data)
	if found {
		return n, found
	}
	return nil, false
}

// LevelOrderPrint 层序遍历 打印二叉树
func (T *BTNode) LevelOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer
	var queue []*BTNode
	queue = append(queue, T)

	for len(queue) > 0 {
		tmp := queue[0]
		queue[0] = nil
		queue = queue[1:]
		ret = append(ret, tmp.Data)
		if tmp.Lchild != nil {
			queue = append(queue, tmp.Lchild)
		}
		if tmp.Rchild != nil {
			queue = append(queue, tmp.Rchild)
		}
	}
	return ret
}

// -----------------------------------
// 非递归的基本思路是模拟递归调用的栈执行顺序
// -----------------------------------

// NRPreOrderPrint 先序遍历非递归
func (T *BTNode) NRPreOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer
	var queue []*BTNode

	tmp := T

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			ret = append(ret, tmp.Data)
			queue = append(queue, tmp)
			tmp = tmp.Lchild
		}

		tmp = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]

		tmp = tmp.Rchild
	}
	return ret
}

// PreOrderPrint 先序遍历
func (T *BTNode) PreOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer

	ret = append(ret, T.Data)
	ret = append(ret, T.Lchild.PreOrderPrint()...)
	ret = append(ret, T.Rchild.PreOrderPrint()...)

	return ret
}

// NRMidOrderPrint 中序遍历非递归
func (T *BTNode) NRMidOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer
	var queue []*BTNode

	tmp := T

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			queue = append(queue, tmp)
			tmp = tmp.Lchild
		}

		tmp = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]

		ret = append(ret, tmp.Data)

		tmp = tmp.Rchild
	}
	return ret
}

// MidOrderPrint 中序遍历
func (T *BTNode) MidOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer

	ret = append(ret, T.Lchild.MidOrderPrint()...)
	ret = append(ret, T.Data)
	ret = append(ret, T.Rchild.MidOrderPrint()...)

	return ret
}

// NRPostOrderPrint 后序遍历非递归
func (T *BTNode) NRPostOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer
	var queue []*nrTmpNode

	tmp := T

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			queue = append(queue, &nrTmpNode{node: tmp, visitNum: 0})
			tmp = tmp.Lchild
		}
		if queue[len(queue)-1].visitNum == 0 {
			queue[len(queue)-1].visitNum = 1
			tmp = queue[len(queue)-1].node.Rchild
		} else {
			nrTmp := queue[len(queue)-1]
			queue[len(queue)-1] = nil
			queue = queue[:len(queue)-1]
			ret = append(ret, nrTmp.node.Data)
			tmp = nil
		}
	}
	return ret
}

// PostOrderPrint 后序遍历
func (T *BTNode) PostOrderPrint() []BTSortDataer {
	if T == nil {
		return nil
	}
	var ret []BTSortDataer

	ret = append(ret, T.Lchild.PostOrderPrint()...)
	ret = append(ret, T.Rchild.PostOrderPrint()...)
	ret = append(ret, T.Data)

	return ret
}
