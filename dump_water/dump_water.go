package dump_water

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// 题目：三个水桶等分8升水
// 来自《算法的乐趣》
// 有3个没有刻度的桶，总容量分别是3,5,8升，其中8升桶里装满了水，求将水等分成4升的倒水步骤
// 思路：状态迁移穷举，将某时刻三个桶的水数量作为一个状态，倒水动作作为一个触发状态迁移的操作
// 从当前状态到下一状态，有n种倒水动作，形成了n叉状态树
// 采用深度优先遍历，可获取到所有倒水的方案
// 优化：状态记忆（类似于动态规划中的子问题缓存）
// 注意点：深度优先遍历可能出现死循环

type Bucket struct {
	cap   int // 桶容量
	water int // 当前水的数量
}

func (b Bucket) String() string {
	return strconv.Itoa(b.water)
}

// LeftCap 桶剩余空间
func (b *Bucket) LeftCap() int {
	return b.cap - b.water
}

// DumpWaterAction 倒水动作
type DumpWaterAction struct {
	from   int // 从哪个桶倒出
	to     int // 倒入到哪个桶
	walter int // 倒多少升水
}

func (da DumpWaterAction) String() string {
	return fmt.Sprintf("<%d,%d,%d>", da.from, da.to, da.walter)
}

// DumpWaterState 状态转移结构体
type DumpWaterState struct {
	state  []Bucket // 每个桶的水的数量，比如初始状态[8,0,0]
	action DumpWaterAction
}

func (ds DumpWaterState) String() string {
	return fmt.Sprintf("%v = %v => %v, ", ds.state, ds.action, ds.NextState())
}

// Validate 是否是有效的倒水状态转移
func (ds *DumpWaterState) Validate() bool {
	// 参数错误
	if ds.action.from < 0 || ds.action.to < 0 || ds.action.from >= len(ds.state) || ds.action.to >= len(ds.state) || ds.action.walter <= 0 {
		return false
	}
	// 不能倒入同一个桶
	if ds.action.from == ds.action.to {
		return false
	}
	// 水不够倒出
	if ds.state[ds.action.from].water < ds.action.walter {
		return false
	}
	// 倒出的水装不下
	if ds.state[ds.action.to].LeftCap() < ds.action.walter {
		return false
	}
	return true
}

// NextState 下一个状态
func (ds *DumpWaterState) NextState() []Bucket {
	next := make([]Bucket, len(ds.state))
	copy(next, ds.state)
	next[ds.action.from].water -= ds.action.walter
	next[ds.action.to].water += ds.action.walter
	return next
}

func DumpWater() {
	initState := []Bucket{
		Bucket{cap: 8, water: 8}, Bucket{cap: 5, water: 0}, Bucket{cap: 3, water: 0},
	}
	finalState := []Bucket{
		Bucket{cap: 8, water: 4}, Bucket{cap: 5, water: 4}, Bucket{cap: 3, water: 0},
	}
	projectNum := int32(0)
	doDumpWater(initState, finalState, []DumpWaterState{}, &projectNum)
}

func doDumpWater(state []Bucket, finalState []Bucket, dss []DumpWaterState, projectNum *int32) {
	// 如果到了最终状态，则打印
	if equal(state, finalState) {
		fmt.Printf("equal-%d:\n", atomic.AddInt32(projectNum, 1))
		for i := 0; i < len(dss); i++ {
			fmt.Println(dss[i].state, "-", dss[i].action, "->", dss[i].NextState())
		}
		fmt.Println("-------------------")
		return
	}
	state = copyBuckets(state)
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state); j++ {
			ds := DumpWaterState{
				state: copyBuckets(state),
				action: DumpWaterAction{
					from:   i,
					to:     j,
					walter: min(state[i].water, state[j].LeftCap()),
				},
			}
			if !ds.Validate() {
				continue
			}
			nextState := ds.NextState()
			found := false
			for k := 0; k < len(dss); k++ {
				if equal(dss[k].state, nextState) {
					found = true
					break
				}
			}
			if !found {
				dss = append(dss, ds)
				doDumpWater(nextState, finalState, dss[0:len(dss):len(dss)], projectNum)
			}
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func copyBuckets(ori []Bucket) []Bucket {
	dst := make([]Bucket, len(ori))
	copy(dst, ori)
	return dst
}

func equal(b1 []Bucket, b2 []Bucket) bool {
	if len(b1) == len(b2) {
		for i := 0; i < len(b1); i++ {
			if b1[i].cap != b2[i].cap || b1[i].water != b2[i].water {
				return false
			}
		}
		return true
	}
	return false
}
