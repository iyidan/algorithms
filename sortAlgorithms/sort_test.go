package sortAlgorithms

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

var testCases = [][]int{
	{5, 4, 3, 2, 1},
	{1, 2, 3, 4, 5},
	{1, 3, 2, 4, 5},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{2, 2, 2, 1, 1, 1, 3, 3, 3, 5, 5, 5, 2, 2, 2},
	{5, 3, 2, 4, 1},
	{5, 3, 2, 12, 2, 2, 4, 5, 6, 8, 7},
	{1, 1, 3, 8, 4, 9, 2, 3},
}

func init() {
	for i := 0; i < 100; i++ {
		rndcase := make([]int, 10)
		for i := 0; i < len(rndcase); i++ {
			rndcase[i] = rand.Intn(math.MaxInt16)
		}
		testCases = append(testCases, rndcase)
	}
}

func TestQuickSort(t *testing.T) {

	for k, v := range testCases {
		t.Log(k, ":bf:", v)
		QuickSort(v)
		t.Log(k, ":af:", v)
		if !sort.IntsAreSorted(v) {
			t.Fatal(v, "not sorted")
		}
	}
}
