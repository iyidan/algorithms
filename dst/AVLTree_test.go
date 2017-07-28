package dst

import "testing"

var (
	nodeRoot = &BTNode{Data: BTDStr("R")}
	nodeP    = &BTNode{Data: BTDStr("P")}
	nodePr   = &BTNode{Data: BTDStr("pr")}
	nodeL    = &BTNode{Data: BTDStr("L")}
	nodeLl   = &BTNode{Data: BTDStr("lr")}
	nodeLr   = &BTNode{Data: BTDStr("ll")}

	nodes = []*BTNode{nodeRoot, nodeP, nodePr, nodeL, nodeLl, nodeLr}
)

func resetNodes() {
	for k := range nodes {
		nodes[k].Lchild = nil
		nodes[k].Rchild = nil
	}
}

func TestAVLRotate(t *testing.T) {
	resetNodes()
	nodeRoot.Lchild = nodeP
	nodeP.Lchild = nodeL
	nodeL.Lchild = nodeLl
	nodeL.Rchild = nodeLr
	nodeP.Rchild = nodePr

	oriStr := nodeRoot.PrettyPrint()

	t.Logf("origin:\n%s\n", nodeRoot.PrettyPrint())

	AVLRotateRight(&nodeRoot.Lchild)

	t.Logf("rotate-right:\n%s\n", nodeRoot.PrettyPrint())

	AVLRotateLeft(&nodeRoot.Lchild)

	t.Logf("rotate-left:\n%s\n", nodeRoot.PrettyPrint())

	if oriStr != nodeRoot.PrettyPrint() {
		t.Fatal("rotate to origin failed")
	}
}

func TestAVLInsert(t *testing.T) {
	var avlT *BTNode
	datas := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}

	for i := range datas {
		AVLInsert(&avlT, BTDInt(datas[i]))
	}

	t.Logf("\n%s\n", avlT.PrettyPrint())
}
