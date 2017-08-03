package dst

import "testing"
import "reflect"

var (

	//         10
	//       /  \
	//      6    18
	//    /   \
	//   3     8
	//    \   /
	//     5 7
	sortedList      = []BTSortDataer{BTDInt(3), BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(8), BTDInt(10), BTDInt(18)}
	sortedListDel7  = []BTSortDataer{BTDInt(3), BTDInt(5), BTDInt(6), BTDInt(8), BTDInt(10), BTDInt(18)}
	sortedListDel5  = []BTSortDataer{BTDInt(3), BTDInt(6), BTDInt(7), BTDInt(8), BTDInt(10), BTDInt(18)}
	sortedListDel10 = []BTSortDataer{BTDInt(3), BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(8), BTDInt(18)}
	sortedListDel6  = []BTSortDataer{BTDInt(3), BTDInt(5), BTDInt(7), BTDInt(8), BTDInt(10), BTDInt(18)}
	sortedListDel3  = []BTSortDataer{BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(8), BTDInt(10), BTDInt(18)}
	sortedListDel8  = []BTSortDataer{BTDInt(3), BTDInt(5), BTDInt(6), BTDInt(7), BTDInt(10), BTDInt(18)}

	noSortList = []BTSortDataer{BTDInt(10), BTDInt(6), BTDInt(18), BTDInt(3), BTDInt(8), BTDInt(5), BTDInt(7)}
	testBSTree = MakeBSTree(noSortList)
)

func TestBSTSearch(t *testing.T) {
	parent, ret := testBSTree.BSTSearch(BTDInt(3))
	if ret == nil {
		t.Fatal("search 3 failed")
	}
	if int(ret.Data.(BTDInt)) != 3 {
		t.Fatal("search ret != 3")
	}
	if int(parent.Data.(BTDInt)) != 6 {
		t.Fatal("search ret  parent != 6")
	}

	parent, ret = testBSTree.BSTSearch(BTDInt(10))
	t.Logf("ret: %p, %p, %p\n", parent, ret, testBSTree)
	if ret == nil {
		t.Fatal("search 10 failed")
	}
	if int(ret.Data.(BTDInt)) != 10 {
		t.Fatal("search ret != 10")
	}
	if parent != ret {
		t.Fatal("parent != ret")
	}

	parent, ret = testBSTree.BSTSearch(BTDInt(MaxInt))
	t.Logf("ret: %p, %p, %p\n", parent, ret, testBSTree.Rchild)
	if ret != nil {
		t.Fatal("search not exists value test failed")
	}
	if int(parent.Data.(BTDInt)) != 18 || parent != testBSTree.Rchild {
		t.Fatal("search not exists value, parent node test failed")
	}
}

func TestBSTInsert(t *testing.T) {
	bst := MakeBSTree(noSortList)
	ret := bst.MidOrderPrint()
	if !reflect.DeepEqual(ret, sortedList) {
		t.Fatalf("TestBSTInsert fail: %#v\n", ret)
	}
}

func TestBSTDelete(t *testing.T) {

	bst := MakeBSTree(noSortList)
	t.Log("origin:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	// 删除不存在的节点
	delNode := bst.BSTDelete(BTDInt(MaxInt))
	if delNode != nil {
		t.Fatal("delete not exists node test fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedList) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedList)`)
	}

	bst = MakeBSTree(noSortList)

	// 删除叶子节点
	delNode = bst.BSTDelete(BTDInt(7))
	if delNode == nil {
		t.Fatal("delete node 7 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel7) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel7)`)
	}
	t.Log("del7:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	bst = MakeBSTree(noSortList)

	delNode = bst.BSTDelete(BTDInt(5))
	if delNode == nil {
		t.Fatal("delete node 5 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel5) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel5)`)
	}
	t.Log("del5:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	bst = MakeBSTree(noSortList)

	// 删除有两个孩子的节点
	delNode = bst.BSTDelete(BTDInt(6))
	if delNode == nil {
		t.Fatal("delete node 6 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel6) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel6)`)
	}
	t.Log("del6:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	bst = MakeBSTree(noSortList)

	// 删除一个孩子的节点
	delNode = bst.BSTDelete(BTDInt(3))
	if delNode == nil {
		t.Fatal("delete node 3 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel3) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel3)`)
	}
	t.Log("del3:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	bst = MakeBSTree(noSortList)

	delNode = bst.BSTDelete(BTDInt(8))
	if delNode == nil {
		t.Fatal("delete node 8 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel8) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel8)`)
	}
	t.Log("del8:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())

	bst = MakeBSTree(noSortList)

	// 删除根节点
	delNode = bst.BSTDelete(BTDInt(10))
	if delNode == nil {
		t.Fatal("delete node 10 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel10) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel10)`)
	}
	t.Log("del10:", bst.MidOrderPrint())
	t.Logf("\n%s\n", bst.PrettyPrint())
}
