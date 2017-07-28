package dst

import (
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
	testTree = &BTNode{
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
	}
	oPre   = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("d"), BTDStr("f"), BTDStr("e"), BTDStr("g"), BTDStr("c")}
	oMid   = []BTSortDataer{BTDStr("d"), BTDStr("f"), BTDStr("b"), BTDStr("g"), BTDStr("e"), BTDStr("a"), BTDStr("c")}
	oLevel = []BTSortDataer{BTDStr("a"), BTDStr("b"), BTDStr("c"), BTDStr("d"), BTDStr("e"), BTDStr("f"), BTDStr("g")}
	oPost  = []BTSortDataer{BTDStr("f"), BTDStr("d"), BTDStr("g"), BTDStr("e"), BTDStr("b"), BTDStr("c"), BTDStr("a")}
)

func TestBTPreSearch(t *testing.T) {
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

func TestBTOrderPrint(t *testing.T) {
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
	n := testTree.GetLayers()
	if n != 4 {
		t.Fatal("testTree layers != 4")
	}
}

func TestPrettyPrint(t *testing.T) {

	data := []int{50, 20, 70, 10, 40, 60, 80, 5, 15, 35, 45, 55, 65, 75, 85, 78, 90, 86}
	tree := &BTNode{Data: BTDInt(data[0])}
	for i := 1; i < len(data); i++ {
		tree.BSTInsert(BTDInt(data[i]))
	}
	t.Logf("\n%s\n", tree.PrettyPrint())

	// tree = &BTNode{Data: BTDInt(data[0])}
	// rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 10; i++ {
	// 	tree.BSTInsert(BTDInt(80 + rander.Intn(15)))
	// }

	// t.Logf("\n%s\n", tree.PrettyPrint())
}
