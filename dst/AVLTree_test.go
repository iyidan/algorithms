package dst

import (
	"math/rand"
	"reflect"
	"testing"
)

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

func genDeleteAVLTree(datas []int) *BTNode {
	var avlT *BTNode
	for i := range datas {
		AVLInsert(&avlT, BTDInt(datas[i]))
	}
	return avlT
}

type avldelTestCase struct {
	Data     []int
	DelData  int
	PreOrder []BTSortDataer
	MidOrder []BTSortDataer
}

var (
	avlDelCases = map[string]*avldelTestCase{
		"leftRight": &avldelTestCase{
			Data:     []int{7, 4, 15, 3, 6, 9, 5},
			DelData:  15,
			PreOrder: []BTSortDataer{BTDInt(6), BTDInt(4), BTDInt(3), BTDInt(5), BTDInt(7), BTDInt(9)},
			MidOrder: []BTSortDataer{BTDInt(3), BTDInt(4), BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(9)},
		},
		"rightLeft": &avldelTestCase{
			Data:     []int{8, 4, 24, 2, 6, 15, 27, 5, 10, 20, 30, 23},
			DelData:  4,
			PreOrder: []BTSortDataer{BTDInt(15), BTDInt(8), BTDInt(5), BTDInt(2), BTDInt(6), BTDInt(10), BTDInt(24), BTDInt(20), BTDInt(23), BTDInt(27), BTDInt(30)},
			MidOrder: []BTSortDataer{BTDInt(2), BTDInt(5), BTDInt(6), BTDInt(8), BTDInt(10), BTDInt(15), BTDInt(20), BTDInt(23), BTDInt(24), BTDInt(27), BTDInt(30)},
		},
		"leftLeft-0": &avldelTestCase{
			Data:     []int{7, 4, 15, 2, 5, 9, 1, 3, 6},
			DelData:  15,
			PreOrder: []BTSortDataer{BTDInt(4), BTDInt(2), BTDInt(1), BTDInt(3), BTDInt(7), BTDInt(5), BTDInt(6), BTDInt(9)},
			MidOrder: []BTSortDataer{BTDInt(1), BTDInt(2), BTDInt(3), BTDInt(4), BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(9)},
		},
		"leftLeft-1": &avldelTestCase{
			Data:     []int{7, 4, 15, 3, 6, 9, 2},
			DelData:  15,
			PreOrder: []BTSortDataer{BTDInt(4), BTDInt(3), BTDInt(2), BTDInt(7), BTDInt(6), BTDInt(9)},
			MidOrder: []BTSortDataer{BTDInt(2), BTDInt(3), BTDInt(4), BTDInt(6), BTDInt(7), BTDInt(9)},
		},
		"rightRight-1": &avldelTestCase{
			Data:     []int{7, 4, 15, 6, 9, 18, 16},
			DelData:  4,
			PreOrder: []BTSortDataer{BTDInt(15), BTDInt(7), BTDInt(6), BTDInt(9), BTDInt(18), BTDInt(16)},
			MidOrder: []BTSortDataer{BTDInt(6), BTDInt(7), BTDInt(9), BTDInt(15), BTDInt(16), BTDInt(18)},
		},
		"rightRight-0": &avldelTestCase{
			Data:     []int{7, 4, 15, 6, 9, 18, 8, 10, 17, 19},
			DelData:  4,
			PreOrder: []BTSortDataer{BTDInt(15), BTDInt(7), BTDInt(6), BTDInt(9), BTDInt(8), BTDInt(10), BTDInt(18), BTDInt(17), BTDInt(19)},
			MidOrder: []BTSortDataer{BTDInt(6), BTDInt(7), BTDInt(8), BTDInt(9), BTDInt(10), BTDInt(15), BTDInt(17), BTDInt(18), BTDInt(19)},
		},
	}
)

func TestAVLDelete(t *testing.T) {

	for k, v := range avlDelCases {
		t.Logf("test-%s:\n", k)

		avlT := genDeleteAVLTree(v.Data)
		t.Logf("b:\n%s\n", avlT.PrettyPrint())

		AVLDelete(&avlT, BTDInt(v.DelData))
		t.Logf("a:\n%s\n", avlT.PrettyPrint())

		opre := avlT.PreOrderPrint()
		if !reflect.DeepEqual(opre, v.PreOrder) {
			t.Fatalf("del-case-fail-preorder: %s, r: %v, c: %v", k, opre, v.PreOrder)
		}
		omid := avlT.MidOrderPrint()
		if !reflect.DeepEqual(omid, v.MidOrder) {
			t.Fatalf("del-case-fail-midorder: %s, r: %v, c: %v", k, omid, v.MidOrder)
		}
	}

	var randcase []int
	for k := 0; k < 2000; k++ {
		for j := 0; j < 200+k; j++ {
			tmp := rand.Intn(2000)
			found := false
			for _, v := range randcase {
				if v == tmp {
					found = true
					break
				}
			}
			if !found {
				randcase = append(randcase, tmp)
			}
		}
		avlT := genDeleteAVLTree(randcase)

		//t.Logf("rand-avt:%v, %v\n", avlT.MidOrderPrint(), avlT.LevelOrderPrint())
		//t.Logf("\n%s\n", avlT.PrettyPrint())
		for i := len(randcase) - 1; i >= 0; i-- {
			//t.Logf("del: %d\n", randcase[i])
			//t.Logf("\n%s\n", avlT.PrettyPrint())
			AVLDelete(&avlT, BTDInt(randcase[i]))
			//t.Logf("\n%s\n", avlT.PrettyPrint())
			//t.Logf("del: %5d, %v\n", randcase[i], avlT.MidOrderPrint())
		}
		if avlT != nil {
			t.Fatal("avl-del-fail: not empty tree")
		}
		randcase = randcase[0:0]
	}

}
