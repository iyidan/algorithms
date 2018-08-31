package dynamic_programming

// CutRod 切割钢条最优问题
// 给定一段长度为n英寸的钢条和一个长度-价格对照表，求收益最大的切割方案
// 长度 1 2 3 4 5  6  7  8  9  10
// 价格 1 5 8 9 10 17 17 20 24 30
// ----
// 1. 刻画最优解的结构特征：
//   不考虑重复的方案，n段钢条有2^(n-1)种切割方案（n寸钢条有n-1个切割点，每个切割点可以选择切还是不切，所以是2的(n-1)次方种）
//   假设最优切割方案为：n = i1 + i2 + ... + ik
//   收益：r(n) = p[i1] + p[i2] + ... + p[ik]
// 2. 递归地定义最优解的值：
//   r(n) = max(p[n], r(1) + r(n-1), r(2) + r(n-2), ..., r(n-1) + r(1))
//   p[n] 对应不切割的方案，
//   其他的n-1个参数对应另外n-1中方案（每个切割点都尝试切和不切）上式可以转化为：
//   r(n) = max(p[i] + r(n-i)) (1<=i<=n) （当i=n的时候对应不切割）
//   只要找到这些方案中的最大收益即可：问题的最优解由相关子问题的最优解组合而成
func CutRod(p []int, n int, Rn map[int]int) int {
	return memoizedCutRod(p, n)
}

// 自顶向下
func memoizedCutRod(p []int, n int) int {
	qarr := make([]int, n+1)
	for i := range qarr {
		qarr[i] = minInt
	}
	return memoizedCutRodAux(p, n, qarr)
}

func memoizedCutRodAux(p []int, n int, qarr []int) int {
	if qarr[n] >= 0 { // 带记忆检测是否有
		return qarr[n]
	}
	q := minInt
	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i++ {
			q = max(q, p[i]+memoizedCutRodAux(p, n-i, qarr))
		}
	}
	qarr[n] = q // 存储上
	return q
}

func bottomUpCutRod(p []int, n int) int {
	qarr := make([]int, n+1)
	for j := 1; j <= n; j++ {
		q := minInt
		for i := 1; i <= j; i++ {
			q = max(q, p[i]+qarr[j-i])
		}
		qarr[j] = q
	}
	return qarr[n]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
