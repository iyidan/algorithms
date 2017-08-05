package dst

import (
	"math/rand"
	"testing"
)

func TestRBTreeInsert(t *testing.T) {
	// randcases
	n := 200
	sortedCase := make([]int, 0, n)
	randCase := make([]int, n)
	for i := 0; i < n; i++ {
		sortedCase = append(sortedCase, i*2)
	}
	copy(randCase, sortedCase)
	for i := 0; i < 2000; i++ {
		ShuffleSliceInt(randCase)
		rbt := MakeRBTreeByInt(randCase)
		if err := CheckRedBlackTree(rbt); err != nil {
			t.Log("case:", randCase)
			t.Logf("\n%s\n", rbt.PrettyPrint())
			t.Fatal(err)
		}
		idx := 0
		rbt.NRMidOrder(func(node *BTNode) bool {
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

func TestRBDelete(t *testing.T) {

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
		rbt := MakeRBTreeByInt(randcase)
		if err := CheckRedBlackTree(rbt); err != nil {
			t.Fatal(err)
		}
		ShuffleSliceInt(randcase)
		//t.Logf("rand-avt:%v, %v\n", rbt.MidOrderPrint(), rbt.LevelOrderPrint())
		//t.Logf("\n%s\n", rbt.PrettyPrint())
		for i := len(randcase) - 1; i >= 0; i-- {
			//fmt.Printf("del: %d\n", randcase[i])
			//fmt.Printf("bf:\n%s\n", rbt.PrettyPrint())
			rbt.RBDelete(BTDInt(randcase[i]))
			//fmt.Printf("af:\n%s\n", rbt.PrettyPrint())
			//fmt.Printf("del: %5d, %v\n", randcase[i], rbt.MidOrderPrint())
			if err := CheckRedBlackTree(rbt); err != nil {
				t.Fatal(err)
			}
		}
		if rbt.root != nil {
			t.Fatal("rbt-del-fail: not empty tree")
		}
		randcase = randcase[0:0]
	}

	// randcases
	sortedCase := make([]int, 0, 50)
	randCase := make([]int, 50)
	for i := 0; i < 50; i++ {
		sortedCase = append(sortedCase, i*2)
	}
	copy(randCase, sortedCase)
	for i := 0; i < 10; i++ {
		ShuffleSliceInt(randCase)
		rbt := MakeRBTreeByInt(randCase)
		if err := CheckRedBlackTree(rbt); err != nil {
			t.Fatal(err)
		}
		for j := 0; j < len(sortedCase); j++ {
			node := rbt.RBDelete(BTDInt(sortedCase[j]))
			if node.Data.Compare(BTDInt(sortedCase[j])) != 0 {
				t.Fatal(`delete fail:`, node, sortedCase[j])
			}
			if err := CheckRedBlackTree(rbt); err != nil {
				t.Fatal(err)
			}
		}
	}

}

// o(logn) 可以根据N来绘图
func BenchmarkBRDelete(b *testing.B) {
	b.StopTimer()
	datas := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		datas = append(datas, i*2)
	}
	rbt := MakeRBTreeByInt(datas)
	ShuffleSliceInt(datas)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rbt.RBDelete(BTDInt(datas[i]))
	}
}

// o(logn) 可以根据N来绘图
func BenchmarkBRInsert(b *testing.B) {
	b.StopTimer()
	datas := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		datas = append(datas, i*2)
	}
	ShuffleSliceInt(datas)
	rbt := &BiTree{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rbt.RBInsert(BTDInt(datas[i]))
	}
}
