package dst

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// BiTree 二叉树，关联root节点和其他信息
type BiTree struct {
	root *BTNode
}

// BTNode 二叉树节点
type BTNode struct {
	Data BTSortDataer
	// 平衡因子（balance factor） AVL树用到
	// color 红黑树用到
	bf     int8
	Lchild *BTNode
	Rchild *BTNode
	Parent *BTNode
}

// TFunc 遍历函数签名
// 返回false退出遍历
type TFunc func(*BTNode) bool

// 非递归后序遍历暂存结构
type nrTmpNode struct {
	node     *BTNode
	visitNum int // 0:第一次访问，1：第二次访问
}

// PreSearch 先序查找指定值是否在树中，
// 如果找到 返回值所对应的节点
func (T *BiTree) PreSearch(data BTSortDataer) (*BTNode, bool) {
	return T.root.preSearch(data)
}

// LevelOrder 层序遍历
func (T *BiTree) LevelOrder(f TFunc) {
	T.root.levelOrder(f)
}

// LevelOrderPrint 层序遍历打印
func (T *BiTree) LevelOrderPrint() []BTSortDataer {
	return T.root.levelOrderPrint()
}

// NRPreOrder 先序遍历非递归
func (T *BiTree) NRPreOrder(f TFunc) {
	T.root.nrPreOrder(f)
}

// NRPreOrderPrint 先序遍历非递归打印
func (T *BiTree) NRPreOrderPrint() []BTSortDataer {
	return T.root.nrPreOrderPrint()
}

// PreOrder 先序遍历
func (T *BiTree) PreOrder(f TFunc) {
	T.root.preOrder(f)
}

// PreOrderPrint 先序遍历打印
func (T *BiTree) PreOrderPrint() []BTSortDataer {
	return T.root.preOrderPrint()
}

// NRMidOrder 中序遍历非递归
func (T *BiTree) NRMidOrder(f TFunc) {
	T.root.nrMidOrder(f)
}

// NRMidOrderPrint 中序遍历非递归
func (T *BiTree) NRMidOrderPrint() []BTSortDataer {
	return T.root.nrMidOrderPrint()
}

// MidOrder 中序遍历
func (T *BiTree) MidOrder(f TFunc) {
	T.root.midOrder(f)
}

// MidOrderPrint 中序遍历
func (T *BiTree) MidOrderPrint() []BTSortDataer {
	return T.root.midOrderPrint()
}

// NRPostOrder 后序遍历非递归
func (T *BiTree) NRPostOrder(f TFunc) {
	T.root.nrPostOrder(f)
}

// NRPostOrderPrint 后序遍历非递归
func (T *BiTree) NRPostOrderPrint() []BTSortDataer {
	return T.root.nrPostOrderPrint()
}

// PostOrder 后序遍历
func (T *BiTree) PostOrder(f TFunc) {
	T.root.postOrder(f)
}

// PostOrderPrint 后序遍历返回
func (T *BiTree) PostOrderPrint() []BTSortDataer {
	return T.root.postOrderPrint()
}

// GetLayers 获取二叉树的层数
func (T *BiTree) GetLayers() int {
	return T.root.getLayers()
}

// TransPlant 将 u从树中删除，并用v来代替 v是u的子树上的节点
func (T *BiTree) TransPlant(u, v *BTNode) {
	if u == nil {
		return
	}
	if u.Parent == nil {
		T.root = v
	} else if u == u.Parent.Lchild {
		u.Parent.Lchild = v
	} else {
		u.Parent.Rchild = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// PrettyPrint 按层序遍历打印，每行一层
func (T *BiTree) PrettyPrint() string {
	return T.root.prettyPrint()
}

func (node *BTNode) preSearch(data BTSortDataer) (*BTNode, bool) {
	if node == nil {
		return nil, false
	}
	if node.Data.Compare(data) == 0 {
		return node, true
	}
	n, found := node.Lchild.preSearch(data)
	if found {
		return n, found
	}
	n, found = node.Rchild.preSearch(data)
	if found {
		return n, found
	}
	return nil, false
}

func (node *BTNode) levelOrder(f TFunc) {
	if node == nil {
		return
	}
	var queue []*BTNode
	queue = append(queue, node)

	for len(queue) > 0 {
		tmp := queue[0]
		queue[0] = nil
		queue = queue[1:]
		if !f(tmp) {
			return
		}
		if tmp.Lchild != nil {
			queue = append(queue, tmp.Lchild)
		}
		if tmp.Rchild != nil {
			queue = append(queue, tmp.Rchild)
		}
	}
}

func (node *BTNode) levelOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.levelOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) nrPreOrder(f TFunc) {
	if node == nil {
		return
	}
	var queue []*BTNode
	tmp := node

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			if !f(tmp) {
				return
			}
			queue = append(queue, tmp)
			tmp = tmp.Lchild
		}

		tmp = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]

		tmp = tmp.Rchild
	}
}

func (node *BTNode) nrPreOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.nrPreOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) preOrder(f TFunc) bool {
	if node == nil {
		return true
	}
	if !f(node) {
		return false
	}
	if !node.Lchild.preOrder(f) {
		return false
	}
	if !node.Rchild.preOrder(f) {
		return false
	}
	return true
}

func (node *BTNode) preOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.preOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) nrMidOrder(f TFunc) {
	if node == nil {
		return
	}
	var queue []*BTNode
	tmp := node

	for tmp != nil || len(queue) > 0 {
		for tmp != nil {
			queue = append(queue, tmp)
			tmp = tmp.Lchild
		}

		tmp = queue[len(queue)-1]
		queue[len(queue)-1] = nil
		queue = queue[:len(queue)-1]

		if !f(tmp) {
			return
		}

		tmp = tmp.Rchild
	}
}

func (node *BTNode) nrMidOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.nrMidOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) midOrder(f TFunc) bool {
	if node == nil {
		return true
	}

	if !node.Lchild.midOrder(f) {
		return false
	}
	if !f(node) {
		return false
	}
	if !node.Rchild.midOrder(f) {
		return false
	}
	return true
}

func (node *BTNode) midOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.midOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) nrPostOrder(f TFunc) {
	if node == nil {
		return
	}
	var queue []*nrTmpNode

	tmp := node

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
			if !f(nrTmp.node) {
				return
			}
			tmp = nil
		}
	}
}

func (node *BTNode) nrPostOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.nrPostOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) postOrder(f TFunc) bool {
	if node == nil {
		return true
	}
	if !node.Lchild.postOrder(f) {
		return false
	}
	if !node.Rchild.postOrder(f) {
		return false
	}
	if !f(node) {
		return false
	}
	return true
}

func (node *BTNode) postOrderPrint() []BTSortDataer {
	var ret []BTSortDataer
	node.postOrder(func(node *BTNode) bool {
		ret = append(ret, node.Data)
		return true
	})
	return ret
}

func (node *BTNode) getLayers() int {
	if node == nil {
		return 0
	}
	leftLayers := 1 + node.Lchild.getLayers()
	rightLayers := 1 + node.Rchild.getLayers()
	if leftLayers > rightLayers {
		return leftLayers
	}
	return rightLayers
}

func freeNode(node *BTNode) {
	if node == nil {
		return
	}
	node.Parent = nil
	node.Lchild = nil
	node.Rchild = nil
	node.bf = 0
}

// btRotateRight 右旋P节点
// 将P节点的左孩子变成L的右孩子
// 将L的右孩子变为P
// 将L节点变成根节点
// 旋转不影响中序遍历顺序
// |         /     |       /       |
// |        P      |      L        |
// |      /   \    |    /   \      |
// |     L     pr  |  ll     P     |
// |   /   \       |       /   \   |
// |  ll    lr     |     lr     pr |
//
// 此处传递 **BTNode是为了方便改树中指向p的父亲节点的指针（Lchild 或者Rchild）
func btRotateRight(T *BiTree, p *BTNode) {
	if p == nil || p.Lchild == nil {
		return
	}
	L := p.Lchild
	oldP := p.Parent

	// 维护parent指针
	L.Parent = oldP
	p.Parent = L
	if L.Rchild != nil {
		L.Rchild.Parent = p
	}

	p.Lchild = L.Rchild
	L.Rchild = p
	if oldP != nil {
		if oldP.Lchild == p {
			oldP.Lchild = L
		} else {
			oldP.Rchild = L
		}
	} else {
		T.root = L
	}
}

// btRotateLeft 左旋P节点
// 将P节点的右孩子变成R的左孩子
// 将R的左孩子变为P
// 将R节点变成根节点
// 旋转不影响中序遍历顺序
// |       /       |         /     |
// |      P        |        R      |
// |    /   \      |      /   \    |
// |  pl     R     |     P     rr  |
// |       /   \   |   /   \       |
// |     rl     rr |  pl    rl     |
//
func btRotateLeft(T *BiTree, p *BTNode) {
	if p == nil || p.Rchild == nil {
		return
	}
	R := p.Rchild
	oldP := p.Parent

	// 维护parent指针
	R.Parent = oldP
	p.Parent = R
	if R.Lchild != nil {
		R.Lchild.Parent = p
	}

	p.Rchild = R.Lchild
	R.Lchild = p
	if oldP != nil {
		if oldP.Lchild == p {
			oldP.Lchild = R
		} else {
			oldP.Rchild = R
		}
	} else {
		T.root = R
	}
}

func (node *BTNode) prettyPrint() string {
	if node == nil {
		return "<nil-tree>"
	}
	var buf bytes.Buffer

	maxLen := 1
	datas := node.nrMidOrderPrint()
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
	layers := node.getLayers()
	sp := 1
	smax := intPow(2, layers-1)*maxLen + (intPow(2, layers-1)-1)*sp
	prettyPrintLevel(&buf, []*BTNode{node}, 1, layers, maxLen, sp, smax)

	return buf.String()
}

// 递归的层序打印树
// sp datalen 必须是奇数
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

// CheckBiTree 检查一个树是不是正常的
// 检查parent指针
func CheckBiTree(T *BiTree) error {
	errorStr := make([]string, 0, 10)
	T.PreOrder(func(node *BTNode) bool {
		if node.Lchild != nil && node.Lchild.Parent != node {
			errorStr = append(errorStr, fmt.Sprintf("CheckBiTree: node.Lchild.Parent != node, node:%v", node.Data))
			if len(errorStr) > 10 {
				return false
			}
		}
		if node.Rchild != nil && node.Rchild.Parent != node {
			errorStr = append(errorStr, fmt.Sprintf("CheckBiTree: node.Rchild.Parent != node, node:%v", node.Data))
			if len(errorStr) > 10 {
				return false
			}
		}
		return true
	})

	if len(errorStr) > 0 {
		return errors.New(strings.Join(errorStr, "\n"))
	}
	return nil
}
