package miniSpanTree

import (
	"fmt"

	"github.com/iyidan/algorithms/dst"
)

// PRIM 普里姆最小生成树算法（mini spanning tree）
// -------
// 步骤
// 从单一顶点开始，普里姆算法按照以下步骤逐步扩大树中所含顶点的数目，直到遍及连通图的所有顶点。
// 输入：一个加权连通图，其中顶点集合为V，边集合为E；
// 初始化：Vnew = {x}，其中x为集合V中的任一节点（起始点），Enew = {}；
// 重复下列操作，直到Vnew = V：
//    1. 在集合E中选取权值最小的边（u, v），其中u为集合Vnew中的元素，而v则是V中没有加入Vnew的顶点
// 如果存在有多条满足前述条件即具有相同权值的边，则可任意选取其中之一；
//    2. 将v加入集合Vnew中，将（u, v）加入集合Enew中；
// 输出：使用集合Vnew和Enew来描述所得到的最小生成树。
// -------
// 证明
// 设prim生成的树为G0
// 假设存在Gmin使得cost(Gmin)<cost(G0)
// 则在Gmin中存在(u,v)不属于G0
// 将(u,v)加入G0中可得一个环，且(u,v)不是该环的最长边
// 这与prim每次生成最短边矛盾
// 故假设不成立，得证.
func PRIM(mg *dst.MGraph) {
	numVex := mg.NumVex()
	lowcost := make([]int, numVex)
	adjvex := make([]int, numVex)
	for i := 0; i < numVex; i++ {
		lowcost[i] = mg.Arcs[0][i]
		adjvex[i] = 0
	}

	for i := 1; i < numVex; i++ {
		min := dst.MaxInt
		k := 0
		for j := 1; j < numVex; j++ {
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
		}
		lowcost[k] = 0
		fmt.Printf("[PRIM]edge: (%s[%d], %s[%d])\n", mg.Vexs[adjvex[k]], adjvex[k], mg.Vexs[k], k)
		for j := 1; j < numVex; j++ {
			if lowcost[j] != 0 && mg.Arcs[k][j] < lowcost[j] {
				lowcost[j] = mg.Arcs[k][j]
				adjvex[j] = k
			}
		}
	}

	costs := 0
	for k, v := range adjvex {
		costs += mg.Arcs[k][v]
	}
	fmt.Printf("[PRIM]costs: %#v\n", costs)
}
