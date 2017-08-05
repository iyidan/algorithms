package dst

import (
	"math/rand"
	"reflect"
	"testing"
)

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
)

func TestCheckBSTree(t *testing.T) {
	bst := MakeBSTree(noSortList)

	if err := CheckBSTree(bst); err != nil {
		t.Fatal(err)
	}

	bst.root.Rchild.Data = BTDInt(1)
	if err := CheckBSTree(bst); err == nil {
		t.Fatal("check non-bst tree fail")
	} else {
		t.Log(err)
	}
}

func TestBSTSearch(t *testing.T) {
	bst := MakeBSTree(noSortList)
	node, ok := bst.BSTSearch(BTDInt(3))
	if !ok {
		t.Fatal("search 3 failed")
	}
	if int(node.Data.(BTDInt)) != 3 {
		t.Fatal("search ret != 3")
	}
	if int(node.Parent.Data.(BTDInt)) != 6 {
		t.Fatal("search ret  parent != 6")
	}

	node, ok = bst.BSTSearch(BTDInt(10))
	if !ok {
		t.Fatal("search 10 failed")
	}
	if int(node.Data.(BTDInt)) != 10 {
		t.Fatal("search ret != 10")
	}

	node, ok = bst.BSTSearch(BTDInt(MaxInt))
	if ok {
		t.Fatal("search not exists value test failed")
	}
	if int(node.Data.(BTDInt)) != 18 || node != bst.root.Rchild {
		t.Fatal("search not exists value, parent node test failed")
	}
	t.Logf("\n%s\n", bst.PrettyPrint())
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
	t.Logf("\n%s\n", bst.PrettyPrint())
	if delNode == nil {
		t.Fatal("delete node 6 fail")
	}
	if !reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel6) {
		t.Fatal(`!reflect.DeepEqual(bst.MidOrderPrint(), sortedListDel6)`)
	}
	t.Log("del6:", bst.MidOrderPrint())

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

func TestBSTDeleteRand(t *testing.T) {
	var randcase []BTSortDataer
	for k := 0; k < 2000; k++ {
		for j := 0; j < 100; j++ {
			tmp := rand.Intn(MaxInt)
			found := false
			for _, v := range randcase {
				if BTDInt(tmp) == v {
					found = true
					break
				}
			}
			if !found {
				randcase = append(randcase, BTDInt(tmp))
			}
		}
		bst := MakeBSTree(randcase)
		if err := CheckBSTree(bst); err != nil {
			t.Fatalf("CheckBSTree failed: %v, %s\n", randcase, err)
		}

		ShuffleSliceBTData(randcase)

		for i := len(randcase) - 1; i >= 0; i-- {
			//fmt.Println("del: ", randcase[i], randcase)
			//fmt.Printf("tbefore:\n%s\n", bst.PrettyPrint())
			bst.BSTDelete(randcase[i])
			//fmt.Printf("tafter:\n%s\n", bst.PrettyPrint())
			if err := CheckBSTree(bst); err != nil {
				t.Fatalf("del %v failed: bst not become binary-search-tree: %s\n", randcase[i], err)
			}
		}
		if bst.root != nil {
			t.Fatalf("bst-del-fail: not empty tree\n%s\n", bst.PrettyPrint())
		}
		randcase = randcase[0:0]
	}
}
