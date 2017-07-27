package miniSpanTree

import (
	"fmt"

	"github.com/iyidan/algorithms/dst"
)

// Kruskal 克鲁斯卡尔算法
// 步骤
// 1. 新建图G，G中拥有原图中相同的节点，但没有边
// 2. 将原图中所有的边按权值从小到大排序
// 3. 从权值最小的边开始，如果这条边连接的两个节点于图G中不在同一个连通分量中，则添加这条边到图G中
// 4. 重复3，直至图G中所有的节点都在同一个连通分量中
// 证明
// 1. 这样的步骤保证了选取的每条边都是桥，因此图G构成一个树。
// 2. 为什么这一定是最小生成树呢？关键还是步骤3中对边的选取。算法中总共选取了n-1条边，
// 每条边在选取的当时，都是连接两个不同的连通分量的权值最小的边
// 3. 要证明这条边一定属于最小生成树，可以用反证法：如果这条边不在最小生成树中，
// 它连接的两个连通分量最终还是要连起来的，通过其他的连法，那么另一种连法与这条边一定构成了环，
// 而环中一定有一条权值大于这条边的边，用这条边将其替换掉，图仍旧保持连通，但总权值减小了。
// 也就是说，如果不选取这条边，最后构成的生成树的总权值一定不会是最小的。
func Kruskal(eg *dst.EdgeGraph) {
	eg.SortEdge()

	parent := make([]int, len(eg.Vexs))
	cost := 0

	for i := 0; i < len(eg.Edges); i++ {
		n := disjointSetFind(parent, eg.Edges[i].Begin)
		m := disjointSetFind(parent, eg.Edges[i].End)
		if n != m {
			parent[n] = m
			fmt.Printf("[kruskal](%s[%d], %s[%d])\n", eg.Vexs[eg.Edges[i].Begin], eg.Edges[i].Begin,
				eg.Vexs[eg.Edges[i].End], eg.Edges[i].End)
			cost += eg.Edges[i].Weight
		}
	}
	fmt.Printf("[kruskal]cost:%d\n", cost)
}

// 回路检测，参考并查集
func disjointSetFind(parent []int, x int) int {
	for parent[x] > 0 {
		x = parent[x]
	}
	return x
}
