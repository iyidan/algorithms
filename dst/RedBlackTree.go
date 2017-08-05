package dst

import (
	"fmt"
)

const (
	// RBRED color red
	RBRED int8 = 0
	// RBBLACK color black
	RBBLACK int8 = 1
)

// MakeRBTree 制造一棵红黑树，用给定的数据
func MakeRBTree(datas []BTSortDataer) *BiTree {
	T := &BiTree{}
	for i := 0; i < len(datas); i++ {
		T.RBInsert(datas[i])
	}
	return T
}

// MakeRBTreeByInt 制造一棵红黑树 用给定的int数组
func MakeRBTreeByInt(datas []int) *BiTree {
	T := &BiTree{}
	for i := 0; i < len(datas); i++ {
		T.RBInsert(BTDInt(datas[i]))
	}
	return T
}

// RBInsert 红黑树插入一个节点
func (T *BiTree) RBInsert(data BTSortDataer) bool {
	if T.root == nil {
		T.root = &BTNode{Data: data, bf: RBBLACK}
		return true
	}
	parent, ok := T.BSTSearch(data)
	if ok {
		return false
	}

	z := &BTNode{Data: data, bf: RBRED, Parent: parent}
	if parent.Data.Compare(data) == 1 {
		parent.Lchild = z
	} else {
		parent.Rchild = z
	}
	rbInsertFixUp(T, z)
	return true
}

// rbInsertFixUp 修复插入的节点不一致属性
// 插入的是红色节点 有以下几种情况
// 1. 插入节点的父节点是黑色的，属性无冲突，跳过
// 2. 插入节点的父节点红色（爷爷节点必定是黑色）：
//		a. 父节点是爷爷节点的左孩子
//			a) 父节点的右兄弟是红色的
//				此时将父节点和父节点的右孩子变为黑色，将爷爷节点变为红色，将z节点上移
//				这种变换不会改变通过爷爷节点路径的黑高
//			b) 父节点的右兄弟是黑色的
//				如果z是右孩子，将z指向父亲节点，并左旋，此时就变成了一种情况：
//			c) 上述 a/b变换后，只有一种情况了：
// 				z是父节点的左孩子，z和父节点都是红色，父节点的右兄弟是黑色，爷爷节点是黑色
// 				此时将爷爷节点染红，将父节点染黑，并以爷爷节点为根做右旋
//		b. 父节点是爷爷节点的右孩子
//			与a.的情况相似，只是左和右进行调换
func rbInsertFixUp(T *BiTree, z *BTNode) {
	for z.Parent != nil && z.Parent.bf == RBRED {
		if z.Parent == z.Parent.Parent.Lchild {
			y := z.Parent.Parent.Rchild
			if y != nil && y.bf == RBRED {
				z.Parent.bf = RBBLACK
				y.bf = RBBLACK
				z.Parent.Parent.bf = RBRED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Rchild {
					z = z.Parent
					btRotateLeft(T, z)
				}
				if z.Parent.Parent != nil {
					z.Parent.bf = RBBLACK
					z.Parent.Parent.bf = RBRED
					btRotateRight(T, z.Parent.Parent)
				}
			}
		} else {
			y := z.Parent.Parent.Lchild
			if y != nil && y.bf == RBRED {
				z.Parent.bf = RBBLACK
				y.bf = RBBLACK
				z.Parent.Parent.bf = RBRED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Lchild {
					z = z.Parent
					btRotateRight(T, z)
				}
				if z.Parent.Parent != nil {
					z.Parent.bf = RBBLACK
					z.Parent.Parent.bf = RBRED
					btRotateLeft(T, z.Parent.Parent)
				}
			}
		}
	}
	T.root.bf = RBBLACK
}

// RBDelete 从树中删除一个节点
// 如果未找到返回nil
// 和常规二叉搜索树删除类似，最后如果删除的实际节点是黑色的，需要修复
// z:原本要删除的节点
// y:实际删除或移动的节点
// x:接替y位置的节点（y的孩子）
func (T *BiTree) RBDelete(data BTSortDataer) *BTNode {
	z, ok := T.BSTSearch(data)
	if !ok {
		return nil
	}

	var x *BTNode
	y := z
	yOriColor := y.bf
	if z.Lchild == nil {
		x = z.Rchild
		T.TransPlant(z, z.Rchild)
	} else if z.Rchild == nil {
		x = z.Lchild
		T.TransPlant(z, z.Lchild)
	} else {
		y := z.Rchild
		for y.Lchild != nil {
			y = y.Lchild
		}
		yOriColor = y.bf
		x = y.Rchild
		if x == nil {
			x = &BTNode{Data: nil, bf: RBBLACK, Parent: y}
		}
		if y.Parent != z {
			x.Parent = y.Parent
			T.TransPlant(y, y.Rchild)
			y.Rchild = z.Rchild
			y.Rchild.Parent = y
		}
		T.TransPlant(z, y)
		y.Lchild = z.Lchild
		y.Lchild.Parent = y
		y.bf = z.bf
	}
	if x == nil {
		x = &BTNode{Data: nil, bf: RBBLACK, Parent: y.Parent}
	}
	if yOriColor == RBBLACK && T.root != nil {
		rbDeleteFixUp(T, x)
	}
	freeNode(z)
	return z
}

// rbDeleteFixUp 删除操作的修复，到这里 说明删除的节点y是黑色的
// 可以想象成，y的黑色附加在x上了（因为y删除或者移动的时候，只有一个孩子x，y的颜色属性下降到x）
// 按此想象：x有黑黑或红黑性质
// 此时分几种情况：
// 1. x是红黑：
//		此种情况直接将x改成黑色（原来的x是红色）即可
// 2. x是黑黑，要想办法将x的这层黑色去掉：
//		a. x是左孩子
//			1)x的右兄弟w如果是红色的，x父节点->红， w改为黑，以x父节点为根左旋
//				（w颜色摇摆到左）
//			x的右兄弟w如果是黑色：
//				2) w的两个孩子都是黑色：将x的一层黑色和w的黑色合并成一层，交个父节点，x上移（x是红黑或者黑黑）
//					如果x是红黑就退出循环，否则继续
//					（x、w黑色合并上移，x升了一级）
//				3) w的左孩子是红色：交换w和其左孩子的颜色（w红，wl黑），以w为根右旋，变成下一种情况
//					（wl把d挤下去，消耗了一层红色，变黑了）
//				4) w的右孩子是红色：
//					将w变成x父节点颜色，将x父节点变成黑色，将w的右孩子变成黑色
//					以x父节点为根左旋
//					（x把多于的黑丢给父节点，父节点把自己的丢给了右孩子，右孩子把自己的再丢给右孩子，右孩子恰好是红色，终止）
//		b. x是右孩子，同左孩子分析
func rbDeleteFixUp(T *BiTree, x *BTNode) {
	for x != T.root && x.bf == RBBLACK {
		if (x.Data == nil && x.Parent.Lchild == nil) || x == x.Parent.Lchild {
			w := x.Parent.Rchild
			if w == nil {
				x = x.Parent
				continue
			}
			if w.bf == RBRED {
				w.bf = RBBLACK
				x.Parent.bf = RBRED
				btRotateLeft(T, x.Parent)
				w = x.Parent.Rchild
			}
			if w == nil || (w.Lchild == nil || w.Lchild.bf == RBBLACK) && (w.Rchild == nil || w.Rchild.bf == RBBLACK) {
				if w != nil {
					w.bf = RBRED
				}
				x = x.Parent
			} else {
				if w.Lchild != nil && w.Lchild.bf == RBRED {
					w.Lchild.bf = RBBLACK
					w.bf = RBRED
					btRotateRight(T, w)
					w = x.Parent.Rchild
				}
				w.bf = x.Parent.bf
				x.Parent.bf = RBBLACK
				if w.Rchild != nil {
					w.Rchild.bf = RBBLACK
				}
				btRotateLeft(T, x.Parent)
				x = T.root
			}
		} else {
			w := x.Parent.Lchild
			if w == nil {
				x = x.Parent
				continue
			}
			if w.bf == RBRED {
				w.bf = RBBLACK
				x.Parent.bf = RBRED
				btRotateRight(T, x.Parent)
				w = x.Parent.Lchild
			}
			if w == nil || (w.Rchild == nil || w.Rchild.bf == RBBLACK) && (w.Lchild == nil || w.Lchild.bf == RBBLACK) {
				if w != nil {
					w.bf = RBRED
				}
				x = x.Parent
			} else {
				if w.Rchild != nil && w.Rchild.bf == RBRED {
					w.Rchild.bf = RBBLACK
					w.bf = RBRED
					btRotateLeft(T, w)
					w = x.Parent.Lchild
				}
				w.bf = x.Parent.bf
				x.Parent.bf = RBBLACK
				if w.Lchild != nil {
					w.Lchild.bf = RBBLACK
				}
				btRotateRight(T, x.Parent)
				x = T.root
			}
		}
	}
	x.bf = RBBLACK
}

// CheckRedBlackTree 检查一棵树是否是红黑树
func CheckRedBlackTree(T *BiTree) error {
	if err := CheckBSTree(T); err != nil {
		return err
	}
	if T.root == nil {
		return nil
	}
	if T.root.bf != RBBLACK {
		return fmt.Errorf("root node color is not BLACK, node=%v", T.root.Data)
	}
	return checkRedBlackTreeInternal(T.root)
}

func checkRedBlackTreeInternal(node *BTNode) error {
	if node == nil {
		return nil
	}

	if node.Lchild != nil {
		if node.Lchild.Data.Compare(node.Data) >= 0 {
			return fmt.Errorf("node's left child egt node, node=%v, left=%v", node.Data, node.Lchild.Data)
		}
		if node.bf == RBRED && node.Lchild.bf == RBRED {
			return fmt.Errorf("node and it's left child color are all red , node=%v, left=%v", node.Data, node.Lchild.Data)
		}
	}
	if node.Rchild != nil {
		if node.Rchild.Data.Compare(node.Data) <= 0 {
			return fmt.Errorf("node's right child elt node, node=%v, right=%v", node.Data, node.Rchild.Data)
		}
		if node.bf == RBRED && node.Rchild.bf == RBRED {
			return fmt.Errorf("node and it's right child color are all red , node=%v, right=%v", node.Data, node.Lchild.Data)
		}
	}

	blackHeights := make(map[string]int)
	node.preOrder(func(pnode *BTNode) bool {
		if pnode.Lchild == nil && pnode.Rchild == nil {
			bh := 0
			tmp := pnode
			for tmp != node.Parent {
				bh += int(tmp.bf)
				tmp = tmp.Parent
			}
			blackHeights[fmt.Sprintf("%v", pnode.Data)] = bh
		}
		return true
	})
	firstbh := -1
	for k, v := range blackHeights {
		if firstbh == -1 {
			firstbh = v
			continue
		}
		if v != firstbh {
			return fmt.Errorf("node to leafnode black num not equal , node=%v, leafnode=%v", node.Data, k)
		}
	}

	ret := checkRedBlackTreeInternal(node.Lchild)
	if ret != nil {
		return ret
	}
	return checkRedBlackTreeInternal(node.Rchild)
}
