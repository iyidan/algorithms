package dst

import (
	"math"
	"reflect"
	"testing"
)

var (

	//         a
	//       /  \
	//      b    c
	//    /   \
	//   d     e
	//    \   /
	//     f g
	oPre   = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("e"), BTDStr("g"), BTDStr("c")}
	oMid   = []BTSortDataer{BTDStr("d"), BTDStr("f"), BTDStr("b"), BTDStr("g"), BTDStr("e"), BTDStr("a"), BTDStr("c")}
	oLevel = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("c"), BTDStr("d"), BTDStr("e"), BTDStr("f"), BTDStr("g")}
	oPost  = []BTSortDataer{BTDStr("f"), BTDStr("d"), BTDStr("g"), BTDStr("e"), BTDStr("b"), BTDStr("c"), BTDStr("a")}

	oPreRotateLeftD  = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("f"), BTDStr("d"), BTDStr("e"), BTDStr("g"), BTDStr("c")}
	oPreRotateRightE = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("g"), BTDStr("e"), BTDStr("c")}

	//           a
	//         /  \
	//        e    c
	//      /
	//     b
	//   /   \
	//   d    g
	//    \
	//     f
	oPreRotateLeftB = []BTSortDataer{BTDStr("a"), BTDStr("e"), BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("g"), BTDStr("c")}

	//           a
	//         /  \
	//        d    c
	//         \
	//          b
	//         /  \
	//        f    e
	//            /
	//           g
	oPreRotateRightB = []BTSortDataer{BTDStr("a"), BTDStr("d"), BTDStr("b"), BTDStr("f"), BTDStr("e"), BTDStr("g"), BTDStr("c")}

	//          c
	//         /
	//        a
	//       /
	//      b
	//    /   \
	//   d     e
	//    \   /
	//     f g
	oPreRotateLeftA = []BTSortDataer{BTDStr("c"), BTDStr("a"), BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("e"), BTDStr("g")}

	//         b
	//       /   \
	//      d      a
	//       \    /  \
	//        f  e    c
	//          /
	//         g
	oPreRotateRightA = []BTSortDataer{BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("a"), BTDStr("e"), BTDStr("g"), BTDStr("c")}
)

func buildTestTree() *BiTree {
	testTree := &BiTree{
		root: &BTNode{
			Data: BTDStr("a"),
			Lchild: &BTNode{
				Data: BTDStr("b"),
				Lchild: &BTNode{
					Data: BTDStr("d"),
					Rchild: &BTNode{
						Data: BTDStr("f"),
					},
				},
				Rchild: &BTNode{
					Data: BTDStr("e"),
					Lchild: &BTNode{
						Data: BTDStr("g"),
					},
				},
			},
			Rchild: &BTNode{
				Data: BTDStr("c"),
			},
		},
	}
	testTree.PreOrder(func(node *BTNode) bool {
		if node.Lchild != nil {
			node.Lchild.Parent = node
		}
		if node.Rchild != nil {
			node.Rchild.Parent = node
		}
		return true
	})
	if err := CheckBiTree(testTree); err != nil {
		panic("CheckBiTree testTree fail:" + err.Error())
	}
	return testTree
}

func TestBTPreSearch(t *testing.T) {
	testTree := buildTestTree()
	ret, ok := testTree.PreSearch(BTDStr("e"))
	if !ok {
		t.Fatal("search e failed")
	}
	if ret.Data != BTDStr("e") {
		t.Fatal("search ret.Data != e")
	}
	_, ok = testTree.PreSearch(BTDStr("ff"))
	if ok {
		t.Fatal("found not exists data: ff!")
	}
}

//         a
//       /  \
//      b    c
//    /   \
//   d     e
//    \   /
//     f g
func TestBTOrder(t *testing.T) {
	testTree := buildTestTree()
	// stop at d
	oPreFragment := []string{"a", "b", "d"}
	// stop at b
	oMidFragment := []string{"d", "f", "b"}
	// stop at b
	oPostFragment := []string{"f", "d", "g", "e", "b"}
	// stop at d
	oLevelFragment := []string{"a", "b", "c", "d"}

	var fragment []string

	// pre order
	testTree.PreOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "d" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oPreFragment) {
		t.Fatal(`PreOrder fragment fail`)
	}
	fragment = fragment[0:0]

	testTree.NRPreOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "d" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oPreFragment) {
		t.Fatal(`NRPreOrder fragment fail`)
	}
	fragment = fragment[0:0]

	// mid order
	testTree.MidOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "b" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oMidFragment) {
		t.Fatal(`MidOrder fragment fail`)
	}
	fragment = fragment[0:0]

	testTree.NRMidOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "b" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oMidFragment) {
		t.Fatal(`NRMidOrder fragment fail`)
	}
	fragment = fragment[0:0]

	// post order
	testTree.PostOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "b" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oPostFragment) {
		t.Fatal(`PostOrder fragment fail`)
	}
	fragment = fragment[0:0]

	testTree.NRPostOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "b" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oPostFragment) {
		t.Fatal(`NRPostOrder fragment fail`)
	}
	fragment = fragment[0:0]

	// level order
	testTree.LevelOrder(func(node *BTNode) bool {
		fragment = append(fragment, string(node.Data.(BTDStr)))
		if fragment[len(fragment)-1] == "d" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(fragment, oLevelFragment) {
		t.Fatal(`LevelOrder fragment fail`)
	}
	fragment = fragment[0:0]
}

func TestBTOrderPrint(t *testing.T) {
	testTree := buildTestTree()
	buildTestTree()
	ret := testTree.LevelOrderPrint()
	if !reflect.DeepEqual(ret, oLevel) {
		t.Fatal(`testTree.LevelOrderPrint fail`, ret)
	}
	t.Log("LevelOrderPrint ok", ret)

	ret = testTree.PreOrderPrint()
	if !reflect.DeepEqual(ret, oPre) {
		t.Fatal(`testTree.PreOrderPrint fail`, ret)
	}
	t.Log("PreOrderPrint ok", ret)
	ret = testTree.NRPreOrderPrint()
	if !reflect.DeepEqual(ret, oPre) {
		t.Fatal(`testTree.NRPreOrderPrint fail`, ret)
	}
	t.Log("NRPreOrderPrint ok", ret)

	ret = testTree.MidOrderPrint()
	if !reflect.DeepEqual(ret, oMid) {
		t.Fatal(`testTree.MidOrderPrint fail`, ret)
	}
	t.Log("MidOrderPrint ok", ret)
	ret = testTree.NRMidOrderPrint()
	if !reflect.DeepEqual(ret, oMid) {
		t.Fatal(`testTree.NRMidOrderPrint fail`, ret)
	}
	t.Log("NRMidOrderPrint ok", ret)

	ret = testTree.PostOrderPrint()
	if !reflect.DeepEqual(ret, oPost) {
		t.Fatal(`testTree.PostOrderPrint fail`, ret)
	}
	t.Log("PostOrderPrint ok", ret)
	ret = testTree.NRPostOrderPrint()
	if !reflect.DeepEqual(ret, oPost) {
		t.Fatal(`testTree.NRPostOrderPrint fail`, ret)
	}
	t.Log("NRPostOrderPrint ok", ret)

}

func TestGetLayers(t *testing.T) {
	testTree := buildTestTree()
	n := testTree.GetLayers()
	if n != 4 {
		t.Fatal("testTree layers != 4")
	}
}

func TestPrettyPrint(t *testing.T) {
	testTree := buildTestTree()
	if intPow(2, -1) != 0 {
		t.Fatal(`intPow(2, -1) != 0`)
	}
	for i := 0; i < 30; i++ {
		if intPow(2, i) != int(math.Pow(2, float64(i))) {
			t.Fatal(`intPow(2, i) fail:`, i, intPow(2, i), int(math.Pow(2, float64(i))))
		}
	}

	t.Logf("\n%s\n", testTree.PrettyPrint())
}

func TestBTRotate(t *testing.T) {

	// left b
	testTree := buildTestTree()
	btRotateLeft(testTree, testTree.root.Lchild)
	pre := testTree.PreOrderPrint()
	mid := testTree.MidOrderPrint()
	if !reflect.DeepEqual(pre, oPreRotateLeftB) {
		t.Fatal(`rotate-left-b: fail-pre`)
	}
	if !reflect.DeepEqual(mid, oMid) {
		t.Fatal(`rotate-left-b: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// right b
	testTree = buildTestTree()
	btRotateRight(testTree, testTree.root.Lchild)
	pre = testTree.PreOrderPrint()
	mid = testTree.MidOrderPrint()
	if !reflect.DeepEqual(pre, oPreRotateRightB) {
		t.Fatal(`rotate-right-b: fail-pre`)
	}
	if !reflect.DeepEqual(mid, oMid) {
		t.Fatal(`rotate-right-b: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// left a
	testTree = buildTestTree()
	btRotateLeft(testTree, testTree.root)
	pre = testTree.PreOrderPrint()
	mid = testTree.MidOrderPrint()
	if !reflect.DeepEqual(pre, oPreRotateLeftA) {
		t.Fatal(`rotate-left-a: fail-pre`)
	}
	if !reflect.DeepEqual(mid, oMid) {
		t.Fatal(`rotate-left-a: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// right a
	testTree = buildTestTree()
	btRotateRight(testTree, testTree.root)
	pre = testTree.PreOrderPrint()
	mid = testTree.MidOrderPrint()
	if !reflect.DeepEqual(pre, oPreRotateRightA) {
		t.Fatal(`rotate-right-a: fail-pre`)
	}
	if !reflect.DeepEqual(mid, oMid) {
		t.Fatal(`rotate-right-a: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// left c
	testTree = buildTestTree()
	btRotateLeft(testTree, testTree.root.Rchild)
	if !reflect.DeepEqual(oPre, testTree.PreOrderPrint()) {
		t.Fatal(`rotate-left-c: fail-pre`)
	}
	if !reflect.DeepEqual(oMid, testTree.MidOrderPrint()) {
		t.Fatal(`rotate-left-c: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}
	// right c
	testTree = buildTestTree()
	btRotateRight(testTree, testTree.root.Rchild)
	if !reflect.DeepEqual(oPre, testTree.PreOrderPrint()) {
		t.Fatal(`rotate-right-c: fail-pre`)
	}
	if !reflect.DeepEqual(oMid, testTree.MidOrderPrint()) {
		t.Fatal(`rotate-right-c: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// left d
	testTree = buildTestTree()
	btRotateLeft(testTree, testTree.root.Lchild.Lchild)
	if !reflect.DeepEqual(oPreRotateLeftD, testTree.PreOrderPrint()) {
		t.Fatal(`rotate-left-d: fail-pre`)
	}
	if !reflect.DeepEqual(oMid, testTree.MidOrderPrint()) {
		t.Fatal(`rotate-left-d: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	// right e
	testTree = buildTestTree()
	btRotateRight(testTree, testTree.root.Lchild.Rchild)
	if !reflect.DeepEqual(oPreRotateRightE, testTree.PreOrderPrint()) {
		t.Fatal(`rotate-right-e: fail-pre`)
	}
	if !reflect.DeepEqual(oMid, testTree.MidOrderPrint()) {
		t.Fatal(`rotate-right-e: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	testTree = buildTestTree()
	btRotateLeft(testTree, testTree.root.Lchild)
	btRotateRight(testTree, testTree.root.Lchild)
	if !reflect.DeepEqual(oPre, testTree.PreOrderPrint()) {
		t.Fatal(`rotate-left-right: fail-pre`)
	}
	if !reflect.DeepEqual(oMid, testTree.MidOrderPrint()) {
		t.Fatal(`rotate-left-right: fail-mid`)
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}
}

func TestTransPlant(t *testing.T) {
	testTree := buildTestTree()

	oTransPre := []BTSortDataer{BTDStr("c")}
	testTree.TransPlant(testTree.root, testTree.root.Rchild)
	if !reflect.DeepEqual(oTransPre, testTree.NRPreOrderPrint()) {
		t.Fatal("transplant fail")
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	testTree = buildTestTree()
	oTransPre = []BTSortDataer{BTDStr("a"), BTDStr("e"), BTDStr("g"), BTDStr("c")}
	testTree.TransPlant(testTree.root.Lchild, testTree.root.Lchild.Rchild)
	if !reflect.DeepEqual(oTransPre, testTree.NRPreOrderPrint()) {
		t.Fatal("transplant fail")
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}

	testTree = buildTestTree()
	oTransPre = []BTSortDataer{BTDStr("a"), BTDStr("g"), BTDStr("c")}
	testTree.TransPlant(testTree.root.Lchild, testTree.root.Lchild.Rchild.Lchild)
	if !reflect.DeepEqual(oTransPre, testTree.NRPreOrderPrint()) {
		t.Fatal("transplant fail")
	}
	if err := CheckBiTree(testTree); err != nil {
		t.Fatal("CheckBiTree fail", err)
	}
}
