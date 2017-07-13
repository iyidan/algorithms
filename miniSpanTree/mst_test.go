package miniSpanTree

import (
	"testing"

	"github.com/iyidan/algorithms/dst"
)

var testMGraph = &dst.MGraph{
	Vexs: []string{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"},
	Arcs: [][]int{
		[]int{0, 10, dst.MaxInt, dst.MaxInt, dst.MaxInt, 11, dst.MaxInt, dst.MaxInt, dst.MaxInt},
		[]int{10, 0, 18, dst.MaxInt, dst.MaxInt, dst.MaxInt, 16, dst.MaxInt, 12},
		[]int{dst.MaxInt, 18, 0, 22, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, 8},
		[]int{dst.MaxInt, dst.MaxInt, 22, 0, 20, dst.MaxInt, 24, 16, 21},
		[]int{dst.MaxInt, dst.MaxInt, dst.MaxInt, 20, 0, 26, dst.MaxInt, 7, dst.MaxInt},
		[]int{11, dst.MaxInt, dst.MaxInt, dst.MaxInt, 26, 0, 17, dst.MaxInt, dst.MaxInt},
		[]int{dst.MaxInt, 16, dst.MaxInt, 24, dst.MaxInt, 17, 0, 19, dst.MaxInt},
		[]int{dst.MaxInt, dst.MaxInt, dst.MaxInt, 16, 7, dst.MaxInt, 19, 0, dst.MaxInt},
		[]int{dst.MaxInt, 12, 8, 21, dst.MaxInt, dst.MaxInt, dst.MaxInt, dst.MaxInt, 0},
	},
}

var testEdgeGraph = &dst.EdgeGraph{
	Vexs: []string{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"},
	Edges: []*dst.Edge{
		&dst.Edge{Begin: 4, End: 7, Weight: 7},
		&dst.Edge{Begin: 2, End: 8, Weight: 8},
		&dst.Edge{Begin: 0, End: 1, Weight: 10},
		&dst.Edge{Begin: 0, End: 5, Weight: 11},
		&dst.Edge{Begin: 1, End: 8, Weight: 12},
		&dst.Edge{Begin: 3, End: 7, Weight: 16},
		&dst.Edge{Begin: 1, End: 6, Weight: 16},
		&dst.Edge{Begin: 5, End: 6, Weight: 17},
		&dst.Edge{Begin: 1, End: 2, Weight: 18},
		&dst.Edge{Begin: 6, End: 7, Weight: 19},
		&dst.Edge{Begin: 3, End: 4, Weight: 20},
		&dst.Edge{Begin: 3, End: 8, Weight: 21},
		&dst.Edge{Begin: 2, End: 3, Weight: 22},
		&dst.Edge{Begin: 3, End: 6, Weight: 24},
		&dst.Edge{Begin: 4, End: 5, Weight: 26},
	},
}

func TestPRIM(t *testing.T) {
	PRIM(testMGraph)
}

func TestKruskal(t *testing.T) {
	Kruskal(testEdgeGraph)
}
