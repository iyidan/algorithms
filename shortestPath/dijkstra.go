package shortestPath

import (
	"fmt"

	"github.com/iyidan/algorithms/dst"
)

// DijksRet 迪杰斯卡尔最短路径算法结果
type DijksRet struct {
	Path []int
	Cost []int
}

// DijkstraDst 到某一点
func DijkstraDst(mg *dst.MGraph, vx int, end int) {
	mg.Print()

	ret := Dijkstra(mg, vx)

	fmt.Println("patharc:")
	for k, v := range ret.Path {
		fmt.Printf("(%d, %d)\n", k, v)
	}
	fmt.Println("cost:")
	for k, v := range ret.Cost {
		fmt.Printf("%d -> %d = %d\n", vx, k, v)
	}

	fmt.Printf("path %d ---> %d:\n", vx, end)
	var p []string
	e := end
	p = append(p, mg.Vexs[end])
	for ret.Path[e] != -1 {
		e = ret.Path[e]
		p = append(p, mg.Vexs[e])
	}
	p = append(p, mg.Vexs[0])
	for i := len(p) - 1; i >= 0; i-- {
		fmt.Printf("-> %s ", p[i])
	}
	fmt.Println("")
}

// Dijkstra 迪杰斯卡尔最短路径算法
func Dijkstra(mg *dst.MGraph, vx int) *DijksRet {
	numVex := mg.NumVex()
	shortestPath := make([]int, numVex)
	patharc := make([]int, numVex)
	finded := make([]int, numVex)
	for i := 0; i < numVex; i++ {
		patharc[i] = -1
		shortestPath[i] = mg.Arcs[vx][i]
		finded[i] = -1
	}
	shortestPath[vx] = 0
	finded[vx] = 0

	for i := 1; i < numVex; i++ {
		min := dst.MaxInt
		k := 0
		for j := 0; j < numVex; j++ {
			if finded[j] == -1 && shortestPath[j] < min {
				k = j
				min = shortestPath[j]
			}
		}
		finded[k] = 0
		for j := 0; j < numVex; j++ {
			if finded[j] == -1 && (min+mg.Arcs[k][j]) < shortestPath[j] {
				shortestPath[j] = min + mg.Arcs[k][j]
				patharc[j] = k
			}
		}
	}
	return &DijksRet{
		Path: patharc,
		Cost: shortestPath,
	}
}
