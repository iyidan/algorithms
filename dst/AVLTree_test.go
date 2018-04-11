package dst

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestAVLInsert(t *testing.T) {
	datas := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	oPre := []int{4, 2, 1, 3, 7, 6, 5, 9, 8, 10}
	oMid := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	avlT := MakeAVLTreeByInt(datas)

	var tmp []int
	avlT.NRPreOrder(func(node *BTNode) bool {
		tmp = append(tmp, int(node.Data.(BTDInt)))
		return true
	})
	if !reflect.DeepEqual(tmp, oPre) {
		t.Fatal("pre order fail, tmp:, correct:,", tmp, oPre)
	}
	tmp = tmp[0:0]
	avlT.NRMidOrder(func(node *BTNode) bool {
		tmp = append(tmp, int(node.Data.(BTDInt)))
		return true
	})
	if !reflect.DeepEqual(tmp, oMid) {
		t.Fatal("middle order fail, tmp:, correct:,", tmp, oMid)
	}
	if err := CheckAVLTree(avlT); err != nil {
		t.Fatal(err)
	}
	avlT.root.Lchild.bf = 3
	if err := CheckAVLTree(avlT); err == nil {
		t.Fatal(`CheckAVLTree bad avltree is ok`)
	} else {
		t.Log(err)
	}
	// t.Logf("\n%s\n", avlT.PrettyPrint())

	// randcases
	sortedCase := make([]int, 0, 200)
	randCase := make([]int, 200)
	for i := 0; i < 200; i++ {
		sortedCase = append(sortedCase, i*2)
	}
	copy(randCase, sortedCase)
	for i := 0; i < 1000; i++ {
		ShuffleSliceInt(randCase)
		avlT = MakeAVLTreeByInt(randCase)
		if err := CheckAVLTree(avlT); err != nil {
			t.Fatal(err)
		}
		idx := 0
		avlT.NRMidOrder(func(node *BTNode) bool {
			d := int(node.Data.(BTDInt))
			if d != sortedCase[idx] {
				t.Fatalf("randCase fail: d=%d, sortedCase[%d]=%d, node=%v\n", d, idx, sortedCase[idx], node)
				return false
			}
			idx++
			return true
		})
	}

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
		// t.Logf("test-%s:\n", k)

		avlT := MakeAVLTreeByInt(v.Data)
		// t.Logf("b:\n%s\n", avlT.PrettyPrint())

		avlT.AVLDelete(BTDInt(v.DelData))
		// t.Logf("a:\n%s\n", avlT.PrettyPrint())

		opre := avlT.PreOrderPrint()
		if !reflect.DeepEqual(opre, v.PreOrder) {
			t.Fatalf("del-case-fail-preorder: %s, r: %v, c: %v", k, opre, v.PreOrder)
		}
		omid := avlT.MidOrderPrint()
		if !reflect.DeepEqual(omid, v.MidOrder) {
			t.Fatalf("del-case-fail-midorder: %s, r: %v, c: %v", k, omid, v.MidOrder)
		}

		if err := CheckAVLTree(avlT); err != nil {
			t.Fatal(k, err)
		}
	}

	var randcase []int
	for k := 0; k < 1000; k++ {
		for j := 0; j < 100; j++ {
			tmp := rand.Intn(MaxInt)
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
		avlT := MakeAVLTreeByInt(randcase)
		if err := CheckAVLTree(avlT); err != nil {
			t.Fatal(err)
		}
		ShuffleSliceInt(randcase)
		//t.Logf("rand-avt:%v, %v\n", avlT.MidOrderPrint(), avlT.LevelOrderPrint())
		//t.Logf("\n%s\n", avlT.PrettyPrint())
		for i := len(randcase) - 1; i >= 0; i-- {
			//t.Logf("del: %d\n", randcase[i])
			//t.Logf("bf:\n%s\n", avlT.PrettyPrint())
			avlT.AVLDelete(BTDInt(randcase[i]))
			//t.Logf("af:\n%s\n", avlT.PrettyPrint())
			//t.Logf("del: %5d, %v\n", randcase[i], avlT.MidOrderPrint())
			if err := CheckAVLTree(avlT); err != nil {
				t.Fatal(err)
			}
		}
		if avlT.root != nil {
			t.Fatal("avl-del-fail: not empty tree")
		}
		randcase = randcase[0:0]
	}

}

// o(logn) 可以根据N来绘图
func BenchmarkAVLDelete(b *testing.B) {
	b.StopTimer()
	datas := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		datas = append(datas, i*2)
	}
	avlT := MakeAVLTreeByInt(datas)
	ShuffleSliceInt(datas)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		avlT.AVLDelete(BTDInt(datas[i]))
	}
}

// o(logn) 可以根据N来绘图
func BenchmarkAVLInsert(b *testing.B) {
	b.StopTimer()
	datas := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		datas = append(datas, i*2)
	}
	ShuffleSliceInt(datas)
	avlT := &BiTree{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		avlT.AVLInsert(BTDInt(datas[i]))
	}
}
