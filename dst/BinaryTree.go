package dst

import (
	"bytes"
	"fmt"
	"math"
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

type prettyPrintQNode struct {
	level int
	node  *BTNode
}

func prettyprintInternal(buf *bytes.Buffer, nodes []*BTNode, level, maxLevel, datalen int) {
	if level > maxLevel {
		return
	}

	floor := float64(maxLevel - level)
	endgeLines := int(math.Pow(2, math.Max(float64(floor-1), 0)))
	firstSpaces := int(math.Pow(2, float64(floor))) - 1
	betweenSpaces := int(math.Pow(2, float64(floor+1))) - 1

	//length := float64(datalen)
	// endgeLines := int(((math.Pow(2, floor)-1)*(length+sp) + 1) / 2)
	// firstSpaces := int(((math.Pow(2, floor)-1)*(length+sp) + 1))
	// betweenSpaces := int(((math.Pow(2, floor+2)-1)*(length+sp)+1)/2) - int(datalen) + 1
	// fmt.Println("-----------------------")
	// fmt.Println("nodes:", prettyprintGetNodesDataString(nodes), "maxLevel:", maxLevel)
	// fmt.Println("floor:", floor)
	// fmt.Println("endgeLines:", endgeLines)
	// fmt.Println("firstSpaces:", firstSpaces)
	// fmt.Println("betweenSpaces:", betweenSpaces)
	// fmt.Println("-----------------------")

	var newNodes []*BTNode

	prettyprintWriteSpaces(buf, firstSpaces)

	for _, node := range nodes {
		if node == nil {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			prettyprintWriteSpaces(buf, datalen)
		} else {
			nodeStr := fmt.Sprintf("%v", node.Data)
			buf.WriteString(nodeStr)
			newNodes = append(newNodes, node.Lchild)
			newNodes = append(newNodes, node.Rchild)
		}
		prettyprintWriteSpaces(buf, betweenSpaces)
	}

	buf.WriteString("\n")
	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			prettyprintWriteSpaces(buf, firstSpaces-i)
			if nodes[j] == nil {
				prettyprintWriteSpaces(buf, endgeLines+endgeLines+i)
				continue
			}
			if nodes[j].Lchild != nil {
				buf.WriteString("/")
			} else {
				prettyprintWriteSpaces(buf, 1)
			}
			prettyprintWriteSpaces(buf, 2*i-1)
			if nodes[j].Rchild != nil {
				buf.WriteString("\\")
			} else {
				prettyprintWriteSpaces(buf, 1)
			}
			prettyprintWriteSpaces(buf, endgeLines+endgeLines-i)
		}
		buf.WriteString("\n")
	}

	// fmt.Println("re---------------------")
	// fmt.Printf("%s", buf.String())
	// fmt.Println("-----------------------")
	prettyprintInternal(buf, newNodes, level+1, maxLevel, datalen)
}

func prettyprintGetNodesDataString(nodes []*BTNode) string {
	ret := []string{"["}
	for _, v := range nodes {
		if v == nil {
			ret = append(ret, "<nil>")
		} else {
			ret = append(ret, fmt.Sprintf("%v", v.Data))
		}
	}
	ret = append(ret, "]")
	return strings.Join(ret, " ")
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

	prettyprintInternal(&buf, []*BTNode{T}, 1, T.GetLayers(), 1)

	return buf.String()
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
