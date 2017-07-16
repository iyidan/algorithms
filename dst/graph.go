package dst

import (
	"fmt"
	"sort"
)

// MGraph graph data structure
// 邻接矩阵
type MGraph struct {
	Vexs []string
	Arcs [][]int
}

// NumVex return mg's vertex num
func (mg *MGraph) NumVex() int {
	return len(mg.Vexs)
}

// Print print a mgraph
func (mg *MGraph) Print() {
	maxLen := len(fmt.Sprintf("%d", MaxInt))
	format := fmt.Sprintf("%%%dd ", maxLen)
	fmt.Println("")
	for _, v := range mg.Arcs {
		fmt.Printf("[")
		for i := range v {
			fmt.Printf(format, v[i])
		}
		fmt.Printf("\b]\n")
	}
	fmt.Println("")
}

// EdgeGraph 边集数组
type EdgeGraph struct {
	Vexs  []string
	Edges []*Edge
}

// SortEdge 按照边的权重值从小到大排序
func (eg *EdgeGraph) SortEdge() {
	if len(eg.Edges) <= 1 {
		return
	}
	sort.Slice(eg.Edges, func(i, j int) bool {
		return eg.Edges[i].Weight < eg.Edges[j].Weight
	})
}

// Edge represent graph's edge
type Edge struct {
	Begin  int
	End    int
	Weight int
}
