package dst

// AVLRotateRight 右旋P节点
// 将P节点的左孩子变成L的右孩子
// 将L的右孩子变为P
// 将L节点变成根节点
//
// |         /     |       /       |
// |        P      |      L        |
// |      /   \    |    /   \      |
// |     L     pr  |  ll     P     |
// |   /   \       |       /   \   |
// |  ll    lr     |     lr     pr |
//
// 此处传递 **BTNode是为了方便改树中指向p的父亲节点的指针（Lchild 或者Rchild）
func AVLRotateRight(p **BTNode) {
	if (*p).Lchild == nil {
		return
	}
	L := (*p).Lchild
	(*p).Lchild = L.Rchild
	L.Rchild = *p
	*p = L
}

// AVLRotateLeft 左旋P节点
// 将P节点的右孩子变成R的左孩子
// 将R的左孩子变为P
// 将R节点变成根节点
//
// |       /       |         /     |
// |      P        |        R      |
// |    /   \      |      /   \    |
// |  pl     R     |     P     rr  |
// |       /   \   |   /   \       |
// |     rl     rr |  pl    rl     |
//
// 此处传递 **BTNode是为了方便改树中指向p的父亲节点的指针（Lchild 或者Rchild）
func AVLRotateLeft(p **BTNode) {
	if (*p).Rchild == nil {
		return
	}
	R := (*p).Rchild
	(*p).Rchild = R.Lchild
	R.Lchild = *p
	*p = R
}

// AVLLeftBalance 对T进行以T为根的最小不平衡二叉树做左平衡旋转(左孩子)
func AVLLeftBalance(T **BTNode) {
	L := (*T).Lchild
	switch L.bf {
	// 新节点是插入到L的左子树上，要做右旋
	case 1:
		(*T).bf = 0
		L.bf = 0
		AVLRotateRight(T)
	// 右边高了，新节点是插入到了L的右子数上，要做双旋转
	case -1:
		Lr := L.Rchild
		switch Lr.bf {
		case 1:
			(*T).bf = -1
			L.bf = 0
		case 0:
			(*T).bf = 0
			L.bf = 0
		case -1:
			(*T).bf = 0
			L.bf = 1
		}
		Lr.bf = 0
		AVLRotateLeft(&(*T).Lchild)
		AVLRotateRight(T)
	}
}

// AVLRightBalance 对T进行以T为根的最小不平衡二叉树做左平衡旋转(右孩子)
func AVLRightBalance(T **BTNode) {
	R := (*T).Rchild
	switch R.bf {
	// 新节点是插入到R的右子树上，要做左旋
	case -1:
		(*T).bf = 0
		R.bf = 0
		AVLRotateLeft(T)
	// 左边高了，新节点是插入到了R的左子数上，要做双旋转
	case 1:
		Rl := R.Lchild
		switch Rl.bf {
		case 1:
			(*T).bf = 0
			R.bf = -1
		case 0:
			(*T).bf = 0
			R.bf = 0
		case -1:
			(*T).bf = 1
			R.bf = 0
		}
		Rl.bf = 0
		AVLRotateRight(&(*T).Rchild)
		AVLRotateLeft(T)
	}
}

// AVLInsert 插入一个节点到T指向的AVL树中
// 第一个返回值代表有没有长高，第二个返回值代表有没有插入
func AVLInsert(T **BTNode, data BTSortDataer) (bool, bool) {

	// 插入新节点，树长高（注意是递归调用，此时的T有可能是树中的某个孩子
	if *T == nil {
		*T = &BTNode{Data: data}
		(*T).bf = 0
		return true, true
	}

	var taller, ok bool

	switch (*T).Data.Compare(data) {
	// 如果节点存在了，则不插入
	case 0:
		return false, false
	// 插入到T的左边，此处是一直递归到顶层
	case 1:
		taller, ok = AVLInsert(&(*T).Lchild, data)
		if !ok { // 如果没有插入ok，直接返回
			return taller, ok
		}
		// 如果树长高了
		if taller {
			switch (*T).bf {
			case 0: // 树原本左右等高，现在插入了一个左边节点，则左高
				(*T).bf = 1
				taller = true
			case 1: // 树原本左高，现在插入了一个左边节点，则左更高，需要平衡
				AVLLeftBalance(T)
				taller = false
			case -1: // 树原本右边高，现在插入了一个左边节点，则等高了
				(*T).bf = 0
				taller = false
			}
		}
	// 插入到T的右边，此处是一直递归到顶层
	case -1:
		taller, ok = AVLInsert(&(*T).Rchild, data)
		if !ok { // 如果没有插入ok，直接返回
			return taller, ok
		}
		// 如果长高了，证明右边插入成功
		if taller {
			switch (*T).bf {
			case 0: // 树原本左右等高，现在插入了一个右边节点，则右高
				(*T).bf = -1
				taller = true
			case 1: // 树原本左高，现在插入了一个右边节点，则等高
				(*T).bf = 0
				taller = false
			case -1: // 树原本右边高，现在插入了一个右边节点，需要平衡
				AVLRightBalance(T)
				taller = false
			}
		}
	}

	return taller, ok
}
