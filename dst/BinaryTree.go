package dst

import (
	"bytes"
	"fmt"
	"strings"
)

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
	bf     int // 平衡因子（balance factor） AVL树用到
	Lchild *BTNode
	Rchild *BTNode
	Parent *BTNode
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

type prettyPrintQNode struct {
	level int
	node  *BTNode
}

// datalen, sp must be odd number
func prettyPrintLevel(buf *bytes.Buffer, nodes []*BTNode, k, h, datalen, sp, smax int) {
	if k > h {
		return
	}
	var (
		firstNodeSp int
		firstLinkSp int
		firstLinkUp int
		nodeBetween int

		nextNodes []*BTNode
	)
	firstLinkSp = ((datalen)/2+1)*intPow(2, h-k-1) - 1
	firstLinkUp = (firstLinkSp + 1) * 2
	firstNodeSp = firstLinkUp - datalen/2 - 1
	if firstNodeSp < 0 {
		firstNodeSp = 0
	}
	if firstLinkSp < 0 {
		firstLinkSp = 0
	}
	if firstLinkUp < 0 {
		firstLinkUp = 0
	}

	if k == 1 {
		nodeBetween = 0
	} else {
		nodeBetween = (smax - 2*firstNodeSp - intPow(2, k-1)*datalen) / (intPow(2, k-1) - 1)
	}

	// fmt.Println("datalen:", datalen)
	// fmt.Println("smax, h, k, sp:", smax, h, k, sp)
	// fmt.Println("firstNodeSp:", firstNodeSp)
	// fmt.Println("firstLinkSp:", firstLinkSp)
	// fmt.Println("firstLinkUp:", firstLinkUp)
	// fmt.Println("nodeBetween:", nodeBetween)

	for i := 1; i <= 3; i++ {
		prettyprintWriteSpaces(buf, firstNodeSp)
		for _, node := range nodes {
			if node == nil {
				if i == 1 {
					nextNodes = append(nextNodes, nil, nil)
				}

				//buf.WriteString(strings.Repeat("x", datalen))
				prettyprintWriteSpaces(buf, datalen)
			} else {
				var nodeStr string
				if i == 1 {
					nodeStr = fmt.Sprintf("d:%v", node.Data)
				} else if i == 2 {
					nodeStr = fmt.Sprintf("b:%d", node.bf)
				} else if i == 3 {
					if node.Parent != nil {
						nodeStr = fmt.Sprintf("p:%v", node.Parent.Data)
					} else {
						nodeStr = fmt.Sprintf("p:%v", "#")
					}
				}
				if len(nodeStr) < datalen {
					wt := strings.Repeat(" ", (datalen-len(nodeStr))/2)
					if (datalen-len(nodeStr))%2 == 0 {
						nodeStr = wt + nodeStr + wt
					} else {
						nodeStr = " " + wt + nodeStr + wt
					}
				}
				buf.WriteString(nodeStr)
				//buf.WriteString(strings.Repeat("o", datalen))
				if i == 1 {
					nextNodes = append(nextNodes, node.Lchild)
					nextNodes = append(nextNodes, node.Rchild)
				}
			}
			prettyprintWriteSpaces(buf, nodeBetween)
		}
		buf.WriteString("\n")
	}

	if k < h {
		for i := 0; i < len(nodes); i++ {
			prettyprintWriteSpaces(buf, firstLinkSp)
			if nodes[i] == nil || (nodes[i].Lchild == nil && nodes[i].Rchild == nil) {
				prettyprintWriteSpaces(buf, firstLinkSp*3+1+1+1+1)
				continue
			}

			if nodes[i].Lchild == nil {
				prettyprintWriteSpaces(buf, firstLinkSp+1)
				buf.WriteString("└")
				buf.WriteString(strings.Repeat("─", firstLinkSp))
				buf.WriteString("┐")
			} else if nodes[i].Rchild == nil {
				buf.WriteString("┌")
				buf.WriteString(strings.Repeat("─", firstLinkSp))
				buf.WriteString("┘")
				prettyprintWriteSpaces(buf, firstLinkSp+1)
			} else {
				buf.WriteString("┌")
				buf.WriteString(strings.Repeat("─", firstLinkSp))
				buf.WriteString("┴")
				buf.WriteString(strings.Repeat("─", firstLinkSp))
				buf.WriteString("┐")
			}

			prettyprintWriteSpaces(buf, firstLinkSp+1)
		}
		buf.WriteString("\n")
	}

	prettyPrintLevel(buf, nextNodes, k+1, h, datalen, sp, smax)
}

func prettyprintWriteSpaces(buf *bytes.Buffer, n int) {
	if n <= 0 {
		return
	}
	buf.Write(bytes.Repeat([]byte{' '}, n))
}

// PrettyPrint 按层序遍历打印，每行一层
func (T *BTNode) PrettyPrint() string {
	if T == nil {
		return "<nil-tree>"
	}
	var buf bytes.Buffer

	maxLen := 1
	datas := T.NRMidOrderPrint()
	for _, v := range datas {
		l := len(fmt.Sprintf("%v", v))
		if l > maxLen {
			maxLen = l
		}
	}
	// d:1
	// b:1
	// p:3
	maxLen = maxLen + 2
	if maxLen%2 == 0 {
		maxLen++
	}
	layers := T.GetLayers()
	sp := 1
	smax := intPow(2, layers-1)*maxLen + (intPow(2, layers-1)-1)*sp
	prettyPrintLevel(&buf, []*BTNode{T}, 1, layers, maxLen, sp, smax)

	return buf.String()
}

// LevelOrder 层序遍历
func (T *BTNode) LevelOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
	var queue []*BTNode
	queue = append(queue, T)

	for len(queue) > 0 {
		tmp := queue[0]
		queue[0] = nil
		queue = queue[1:]
		f(tmp)
		if tmp.Lchild != nil {
			queue = append(queue, tmp.Lchild)
		}
		if tmp.Rchild != nil {
			queue = append(queue, tmp.Rchild)
		}
	}
}

// LevelOrderPrint 层序遍历打印
func (T *BTNode) LevelOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.LevelOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// -----------------------------------
// 非递归的基本思路是模拟递归调用的栈执行顺序
// -----------------------------------

// NRPreOrder 先序遍历非递归
func (T *BTNode) NRPreOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
	var queue []*BTNode
	tmp := T

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			f(tmp)
			queue = append(queue, tmp)
			tmp = tmp.Lchild
		}

		tmp = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]

		tmp = tmp.Rchild
	}
}

// NRPreOrderPrint 先序遍历非递归打印
func (T *BTNode) NRPreOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.NRPreOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// PreOrder 先序遍历
func (T *BTNode) PreOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
	f(T)
	T.Lchild.PreOrder(f)
	T.Rchild.PreOrder(f)
}

// PreOrderPrint 先序遍历打印
func (T *BTNode) PreOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.PreOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// NRMidOrder 中序遍历非递归
func (T *BTNode) NRMidOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
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

		f(tmp)

		tmp = tmp.Rchild
	}
}

// NRMidOrderPrint 中序遍历非递归
func (T *BTNode) NRMidOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.NRMidOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// MidOrder 中序遍历
func (T *BTNode) MidOrder(f func(*BTNode)) {
	if T == nil {
		return
	}

	T.Lchild.MidOrder(f)
	f(T)
	T.Rchild.MidOrder(f)
}

// MidOrderPrint 中序遍历
func (T *BTNode) MidOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.MidOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// NRPostOrder 后序遍历非递归
func (T *BTNode) NRPostOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
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
			f(nrTmp.node)
			tmp = nil
		}
	}
}

// NRPostOrderPrint 后序遍历非递归
func (T *BTNode) NRPostOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.NRPostOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

// PostOrder 后序遍历
func (T *BTNode) PostOrder(f func(*BTNode)) {
	if T == nil {
		return
	}
	T.Lchild.PostOrder(f)
	T.Rchild.PostOrder(f)
	f(T)
}

// PostOrderPrint 后序遍历返回
func (T *BTNode) PostOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	T.PostOrder(func(node *BTNode) {
		ret = append(ret, node.Data)
	})
	return ret
}

//GetLayers 获取二叉树的层数
func (T *BTNode) GetLayers() int {
	if T == nil {
		return 0
	}
	leftLayers := 1 + T.Lchild.GetLayers()
	rightLayers := 1 + T.Rchild.GetLayers()
	if leftLayers > rightLayers {
		return leftLayers
	}
	return rightLayers
}

// FreeNode 释放一个node的指针
func FreeNode(node *BTNode) {
	if node == nil {
		return
	}
	node.Parent = nil
	node.Lchild = nil
	node.Rchild = nil
	node.bf = 0
}
