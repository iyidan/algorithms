package dst

import (
	"fmt"
)

// MakeAVLTree 制造一棵avl树，用给定的数据
func MakeAVLTree(datas []BTSortDataer) *BiTree {
	T := &BiTree{}
	for i := 0; i < len(datas); i++ {
		T.AVLInsert(datas[i])
	}
	return T
}

// MakeAVLTreeByInt 制造一棵avl树 用给定的int数组
func MakeAVLTreeByInt(datas []int) *BiTree {
	T := &BiTree{}
	for i := 0; i < len(datas); i++ {
		T.AVLInsert(BTDInt(datas[i]))
	}
	return T
}

// 右旋计算pf值，P为根，L为P的左孩子
// BFP2 = BFP - 1 -BFL
// BFL2 = BFL -1

// 左旋->右旋
// 涉及三个节点的bf：P L LR
// 最终的结果：
// 方便理解，可以想象决定树高的路径，
// P： 右孩子没变，左孩子变成了 LR的右孩子，相当于P的左边去掉了两个节点
// 		如果LR的bf小于0（LR的右子树高），树高决定性路径由LR的右子树承担
//		如果LR的bf大于=0，（LR左子树高或者等高），树高的决定性路径由LR的左子树承担，此时的P要减-BFLR
// 		BFP2  = BFP -2 + (0|-BFLR) （如果BFLR <= 0取0，否则取-BFLR）
// L： 此时L左孩子没变，右孩子变为LR的左孩子
// 		如果LR的bf>=0，树高路径由LR的左孩子决定，相当于L的右边少了一个节点
//		如果LR的bf<0，树高路径由LR的右孩子决定，想当于L的右边少了2个节点
// 		BFL2  = BFL + 1 + (0|-BFLR) (如果BFLR >=0 取0，否则取-BFLR)
// LR：由于决定树高的都是LR的左右子树，变形后的结果是在LR的左右两边分别加了1个节点
//		而且左右子树等高了，因为如果有差异，也会由P的右子树或者L的左子树补齐
// 		BFLR2 = 0

// 左旋计算pf值，P为根，R为P的右孩子
// 此时P为-2，可以想象将P的左增加2个节点才能与右子树等高
// BFP2 = BFP + 1 + (-BFR)
// BFR2 = BFR +1

// 右旋-左旋
// 涉及三个节点的bf：P R RL
// 最终的结果：
// 方便理解，可以想象决定树高的路径，
// P： 左孩子没变，右孩子变成了 RL的左孩子，相当于P的右边去掉了两个节点
// 		如果RL的bf<0（LR的右子树高），树高决定性路径由LR的右子树承担
//		如果RL的bf>=0，（LR左子树高或者等高），树高的决定性路径由LR的左子树承担，此时的P要减-BFLR
// 		BFP2  = BFP +2 + (0|-BFRL) （如果BFRL >= 0取0，<0取-BFRL）
// R： 此时R右孩子没变，左孩子变为RL的右孩子
// 		如果RL的bf<=0，树高路径由RL的右孩子决定，相当于R的左边少了一个节点
//		如果RL的bf>0，树高路径由RL的左孩子决定，想当于R的左边少了2个节点
// 		BFR2  = BFL - 1 + (0|-BFRL) (如果BFRL <=0 取0，>0取-BFRL)
// RL：由于决定树高的都是RL的左右子树，变形后的结果是在RL的左右两边分别加了1个节点
//		而且左右子树等高了，因为如果有差异，也会由P的左子树或者R的右子树补齐
// 		BFRL2 = 0

// avlLeftBalance 对T进行以p为根的最小不平衡二叉树做左平衡旋转(左孩子)
func avlLeftBalance(T *BiTree, p *BTNode) {
	L := p.Lchild
	switch L.bf {
	// 左-左
	case 0:
		fallthrough
	case 1:
		p.bf = p.bf - 1 - L.bf
		L.bf = L.bf - 1
		btRotateRight(T, p)
	// 左-右，要做双旋转
	case -1:
		Lr := L.Rchild
		if Lr.bf >= 0 {
			p.bf = p.bf - 2 - Lr.bf
			L.bf = L.bf + 1
		} else {
			p.bf = p.bf - 2
			L.bf = L.bf + 1 - Lr.bf
		}
		Lr.bf = 0
		btRotateLeft(T, p.Lchild)
		btRotateRight(T, p)
	}
}

// avlRightBalance 对T进行以T为根的最小不平衡二叉树做右平衡旋转(右孩子)
func avlRightBalance(T *BiTree, p *BTNode) {
	R := p.Rchild
	switch R.bf {
	// 右-右
	case 0:
		fallthrough
	case -1:
		p.bf = p.bf + 1 + (-R.bf)
		R.bf = R.bf + 1
		btRotateLeft(T, p)
	// 右-左
	case 1:
		Rl := R.Lchild
		if Rl.bf >= 0 {
			p.bf = p.bf + 2
			R.bf = R.bf - 1 - Rl.bf
		} else {
			p.bf = p.bf + 2 - Rl.bf
			R.bf = R.bf - 1
		}
		Rl.bf = 0
		btRotateRight(T, p.Rchild)
		btRotateLeft(T, p)
	}
}

func avlInsertFixUp(T *BiTree, x *BTNode) {
	taller := true
	for taller && x != nil && x.Parent != nil {
		if x.Parent.Lchild == x {
			switch x.Parent.bf {
			case -1:
				x.Parent.bf = 0
				taller = false
			case 0:
				x.Parent.bf = 1
				x = x.Parent
			case 1:
				x.Parent.bf = 2
				// x.bf = 1/-1的时候树都降低高度，可以退出循环
				if x.bf != 0 {
					taller = false
				}
				avlLeftBalance(T, x.Parent)
				x = x.Parent
			}
		} else {
			switch x.Parent.bf {
			case 1:
				x.Parent.bf = 0
				taller = false
			case 0:
				x.Parent.bf = -1
				x = x.Parent
			case -1:
				x.Parent.bf = -2
				// x.bf = 1/-1的时候树都降低高度，可以退出循环
				if x.bf != 0 {
					taller = false
				}
				avlRightBalance(T, x.Parent)
				x = x.Parent
			}
		}
	}
}

// AVLInsert 插入一个节点到T指向的AVL树中
// 第一个返回值代表有没有长高，第二个返回值代表有没有插入
func (T *BiTree) AVLInsert(data BTSortDataer) bool {
	if T.root == nil {
		T.root = &BTNode{Data: data, bf: 0}
		return true
	}
	parent, ok := T.BSTSearch(data)
	if ok {
		return false
	}

	newNode := &BTNode{Data: data, bf: 0, Parent: parent}
	if parent.Data.Compare(data) == 1 {
		parent.Lchild = newNode
	} else {
		parent.Rchild = newNode
	}

	avlInsertFixUp(T, newNode)
	return true

}

func avlDeleteFixUp(T *BiTree, x *BTNode, isLeft bool) {
	fixup := true
	for fixup {
		// x左右孩子一样高，删掉左边后，整个x子树整体高度不变
		switch x.bf {
		case 0:
			if isLeft {
				x.bf = -1
			} else {
				x.bf = 1
			}
			fixup = false
		case 1:
			if isLeft {
				x.bf = 0
			} else {
				x.bf = 2
				// x.bf = 0树高度不变了，可以退出循环
				if x.Lchild.bf == 0 {
					fixup = false
				}
				avlLeftBalance(T, x)
				x = x.Parent
			}
		case -1:
			if isLeft {
				x.bf = -2
				// x.bf = 0树高度不变了，可以退出循环
				if x.Rchild.bf == 0 {
					fixup = false
				}
				avlRightBalance(T, x)
				x = x.Parent
			} else {
				x.bf = 0
			}
		}
		if fixup {
			if x.Parent == nil {
				fixup = false
			} else if x.Parent.Lchild == x {
				isLeft = true
				x = x.Parent
			} else {
				isLeft = false
				x = x.Parent
			}
		}
	}
}

// AVLDelete 在树中删除一个节点
func (T *BiTree) AVLDelete(data BTSortDataer) *BTNode {
	if T.root == nil {
		return nil
	}
	node, ok := T.BSTSearch(data)
	if !ok {
		return nil
	}

	x := node.Parent // 真正要删除或者移动的节点的父节点
	isLeft := false  // 是不是删除x的左节点

	// 判断删除的情况
	if node.Lchild == nil || node.Rchild == nil {
		tmp := node.Lchild
		if tmp == nil {
			tmp = node.Rchild
		}
		if tmp != nil {
			tmp.Parent = node.Parent
		}
		if node.Parent == nil {
			T.root = tmp
		} else if node.Parent.Lchild == node {
			node.Parent.Lchild = tmp
			isLeft = true
		} else {
			node.Parent.Rchild = tmp
			isLeft = false
		}
	} else {
		// 如果两个孩子都有，找到node的直接后继
		successor := node.Rchild
		for successor.Lchild != nil {
			successor = successor.Lchild
		}
		successor.bf = node.bf // 将node的bf赋值给successor
		if successor.Parent != node {
			x = successor.Parent
			if successor.Parent.Lchild == successor {
				isLeft = true
			} else {
				isLeft = false
			}
			T.TransPlant(successor, successor.Rchild)
			successor.Rchild = node.Rchild
			successor.Rchild.Parent = successor
		} else {
			x = successor
			isLeft = false
		}
		T.TransPlant(node, successor)
		successor.Lchild = node.Lchild
		node.Lchild.Parent = successor
	}

	if x != nil {
		avlDeleteFixUp(T, x, isLeft)
	}

	freeNode(node)
	return node
}

// CheckAVLTree 检查树是否满足avl条件
func CheckAVLTree(T *BiTree) error {
	err := CheckBSTree(T)
	if err != nil {
		return err
	}
	T.LevelOrder(func(node *BTNode) bool {
		layerLeft := 0
		layerRight := 0
		if node.Lchild != nil {
			layerLeft = node.Lchild.getLayers()
		}
		if node.Rchild != nil {
			layerRight = node.Rchild.getLayers()
		}
		bf := int8(layerLeft - layerRight)
		if bf < -1 || bf > 1 {
			err = fmt.Errorf("CheckAVLTree: node balance factor gt 1, node: %v", node)
			return false
		}
		if node.bf != bf {
			err = fmt.Errorf("CheckAVLTree: node.bf != bf, node: %v, bf=%d", node, bf)
			return false
		}
		return true
	})
	return err
}
