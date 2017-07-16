package shortestPath

import (
	"testing"

	"github.com/iyidan/algorithms/dst"
)

var testMGraph = &dst.MGraph{
	Vexs: []string{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"},
	Arcs: [][]int{
		[]int{0, 1, 5, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt},
		[]int{1, 0, 3, 7, 5, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt},
		[]int{5, 3, 0, dst.MaxInt, 1, 7, dst.MaxInt, dst.MaxInt, dst.MaxInt},
		[]int{dst.MaxInt, 7, dst.MaxInt, 0, 2, dst.MaxInt, 3, dst.MaxInt, dst.MaxInt},
		[]int{dst.MaxInt, 5, 1, 2, 0, 3, 6, 9, dst.MaxInt},
		[]int{dst.MaxInt, dst.MaxInt, 7, dst.MaxInt, 3, 0, dst.MaxInt, 5, dst.MaxInt},
		[]int{dst.MaxInt, dst.MaxInt, dst.MaxInt, 3, 6, dst.MaxInt, 0, 2, 7},
		[]int{dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, 9, 5, 2, 0, 4},
		[]int{dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, 7, 4, 0},
	},
}

func TestDijkstra(t *testing.T) {
	DijkstraDst(testMGraph, 0, 8)
}
