package skiplist

import (
	"log"
	"strings"
	"testing"
)

type testStruct struct {
	data string
}

var cf = func(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}
var cfStruct = func(a, b interface{}) int {
	return strings.Compare(a.(testStruct).data, b.(testStruct).data)
}
var cfPointStruct = func(a, b interface{}) int {
	return strings.Compare(a.(*testStruct).data, b.(*testStruct).data)
}

func TestAddStruct(t *testing.T) {
	sl := New(cfStruct)
	sl.Add(testStruct{"t1"}, 10)
	sl.Add(testStruct{"t2"}, 10)
	sl.Add(testStruct{"t3"}, 9)

	log.Printf("%#v", sl)
	log.Println(sl.Score(testStruct{"t1"}))
	sl.Del(testStruct{"t1"})
	log.Printf("%#v", sl)
	log.Println(sl.Len())

	if sl.Len() != 2 || len(sl.scoreMap) != 2 {
		t.Fatal("del struct failed")
	}
}

func TestAddPointStruct(t *testing.T) {
	sl := New(cfPointStruct)

	a := &testStruct{"t1"}
	b := &testStruct{"t2"}
	c := &testStruct{"t3"}

	sl.Add(a, 10)
	sl.Add(b, 10)
	sl.Add(c, 9)

	log.Printf("%#v", sl)
	log.Println(sl.Score(a))
	sl.Del(a)
	log.Printf("%#v", sl)
	log.Println(sl.Len())

	if sl.Len() != 2 || len(sl.scoreMap) != 2 {
		t.Fatal("del struct failed")
	}
}

func TestZslAdd(t *testing.T) {
	sl := New(cf)
	if !sl.Add("t1", 10) {
		t.Fatal("add failed")
	}
	if sl.Add("t1", 10) {
		t.Fatal("add same value failed")
	}
	if sl.length != 1 {
		t.Fatal("add one value, length != 1")
	}
	if !sl.Add("t1", 11) {
		t.Fatal("add same value and not same score: failed")
	}
	if sl.length != 1 {
		t.Fatal("add one value, update it score,  length != 1")
	}

	sl.Add("t2", 1)
	if sl.length != 2 {
		t.Fatal("add two value, length != 2")
	}

}

func TestZslDel(t *testing.T) {
	sl := New(cf)
	sl.Add("t1", 10)
	if sl.Del("t1") != 1 {
		t.Fatal("del t1failed")
	}
	if sl.Del("t1") != 0 {
		t.Fatal("del not exists t1 failed")
	}

	sl.Add("t2", 1)
	if *sl.tail.data != "t2" {
		t.Fatal(`*sl.tail.data != "t2"`)
	}

	sl.Add("t3", 2)
	if *sl.tail.data != "t3" {
		t.Fatal(`*sl.tail.data != "t3"`)
	}
	sl.Add("t4", 3)
	sl.Del("t4")
	if *sl.tail.data != "t3" {
		t.Fatal(`*sl.tail.data != "t3", 2`)
	}

	if sl.Del("t3", "t3", "t2") != 2 {
		t.Fatal(`sl.Del("t3", "t3", "t2") != 2`)
	}
	if sl.length != 0 {
		t.Fatal("del value length != 0")
	}
}

func TestZslRangeByScore_1(t *testing.T) {
	sl := New(cf)
	sl.Add("t1", 1)
	sl.Add("t2", 2)
	sl.Add("t3", 3)
	sl.Add("t4", 4)
	sl.Add("t5", 5)
	sl.Add("t6", 6)
	sl.Add("t7", 7)

	lst := sl.RangeByScore(10, 8)
	if len(lst) != 0 {
		t.Fatal("len(lst) != 0")
	}

	lst = sl.RangeByScore(-1, 100)
	if len(lst) != 7 {
		t.Fatal("len(lst) != 7")
	}
	if lst[0].Data != "t1" {
		t.Fatal(`lst[0].Data != "t1"`)
	}
	if lst[6].Data != "t7" {
		t.Fatal(`lst[6].Data != "t7"`)
	}

	t.Logf("%#v\n", *lst[6])
}

func TestZslRangeByScore_2(t *testing.T) {
	sl := New(cf)
	sl.Add("t1", 1)
	sl.Add("t2", 2)
	sl.Add("t3", 3)
	sl.Add("t4", 4)
	sl.Add("t5", 5)
	sl.Add("t6", 6)
	sl.Add("t7", 7)

	lst := sl.RangeByScore(2, 5)
	if len(lst) != 3 {
		t.Fatal("len(lst) != 3")
	}
	if lst[0].Data != "t2" {
		t.Fatal(`lst[0].Data != "t2"`)
	}
	if lst[2].Data != "t4" {
		t.Fatal(`lst[2].Data != "t4"`)
	}

	t.Logf("%#v\n", *lst[1])
}

func TestZslRangeByScore_3(t *testing.T) {
	sl := New(cf)
	sl.Add("t1_0", 1)
	sl.Add("t1", 1)
	sl.Add("t2", 2)
	sl.Add("t2_1", 2)
	sl.Add("t2_2", 2)
	sl.Add("t3", 3)
	sl.Add("t4_0", 2.5)
	sl.Add("t4", 4)
	sl.Add("t4_1", 4)
	sl.Add("t4_2", 4)
	sl.Add("t5", 5)
	sl.Add("t6", 6)
	sl.Add("t7", 7)

	lst := sl.RangeByScore(1, 5)
	for k, v := range lst {
		t.Logf("lst[%d]=%#v\n", k, v)
	}
	if len(lst) != 10 {
		t.Fatal("len(lst) != 10")
	}
	if lst[0].Data != "t1" {
		t.Fatal(`lst[0].Data != "t1"`)
	}
	if lst[5].Data != "t4_0" {
		t.Fatal(`lst[5].Data != "t4_0"`)
	}
}

func TestZskRank(t *testing.T) {
	sl := New(cf)
	sl.Add("t1", 1)
	sl.Add("t2", 2)
	sl.Add("t3", 3)
	sl.Add("t4", 4)
	sl.Add("t5", 5)
	sl.Add("t6", 6)
	sl.Add("t7", 7)

	if sl.Rank("t1") != 1 {
		t.Logf("%#v", sl.Rank("t1"))
		t.Fatal("rank(t1) != 1")
	}
	if sl.Rank("t4") != 4 {
		t.Fatal("rank(t4) != 4")
	}
	if sl.Rank("t7") != 7 {
		t.Fatal("rank(t7) != 7")
	}

	sl.Add("t2.5", 2.5)
	if sl.Rank("t2.5") != 3 {
		t.Fatal("Add t2.5, rank(t2.5) != 3")
	}

	sl.Del("t2")
	if sl.Rank("t2") != 0 {
		t.Fatal("Del t2, rank(t2) != 0")
	}

	if sl.Rank("t2.5") != 2 {
		t.Fatal("Del t2, rank(t2.5) != 2")
	}

	if sl.Rank("t7") != 7 {
		t.Fatal("Del t2, rank(t7) != 7")
	}
}

func TestRangeByRank(t *testing.T) {
	sl := New(cf)

	lst := sl.RangeByRank(0, 10)
	if len(lst) != 0 {
		t.Fatalf("RangeByRank empty sl failed, len=%d\n", len(lst))
	}

	sl.Add("t1", 1)
	sl.Add("t2", 1)
	sl.Add("t3", 2)
	sl.Add("t4", 4)
	sl.Add("t5", 4)
	sl.Add("t6", 6)
	sl.Add("t7", 7)

	lst = sl.RangeByRank(0, 10)
	for k, v := range lst {
		t.Logf("a, lst[%d]=%#v\n", k, v)
	}
	if len(lst) != 7 {
		t.Fatal("a, len(lst) != 7")
	}

	lst = sl.RangeByRank(1, 2)
	for k, v := range lst {
		t.Logf("b, lst[%d]=%#v\n", k, v)
	}
	if len(lst) != 1 {
		t.Fatal("b, len(lst) != 1")
	}

	lst = sl.RangeByRank(1, 7)

	for k, v := range lst {
		t.Logf("c, lst[%d]=%#v\n", k, v)
	}
	if len(lst) != 6 {
		t.Fatal("c, len(lst) != 1")
	}
}

func TestLen(t *testing.T) {
	sl := New(cf)
	if sl.Len() != 0 {
		t.Fatal("len failed1")
	}
	sl.Add("t1", 1)
	sl.Add("t2", 1)
	sl.Add("t3", 2)

	if sl.Len() != 3 {
		t.Fatal("len failed2")
	}

	sl.Del("t3")

	if sl.Len() != 2 {
		t.Fatal("len failed3")
	}
}

func TestScore(t *testing.T) {
	sl := New(cf)

	score, ok := sl.Score("not exists")
	if ok {
		t.Fatal("not exists")
	}

	sl.Add("t1", -1)
	if score, _ = sl.Score("t1"); score != -1 {
		t.Fatal(`sl.Score("t1") != -1`)
	}
	sl.Add("t2", 0)
	if score, _ = sl.Score("t2"); score != 0 {
		t.Fatal(`sl.Score("t2") != 0`)
	}
	sl.Add("t1", 2)
	if score, _ = sl.Score("t1"); score != 2 {
		t.Fatal(`sl.Score("t1") != 2`)
	}

}

func TestAddMinusScore(t *testing.T) {
	sl := New(cf)

	sl.Add("t1", -2)
	sl.Add("t2", -1)
	sl.Add("t3", 0)
	sl.Add("t4", 1)
	sl.Add("t5", 2)

	if sl.Rank("t5") != 5 {
		t.Fatal(`sl.Rank("t5") != 5`)
	}

	d := sl.RangeByRank(1, 6)
	t.Logf("%#v", d)
}
