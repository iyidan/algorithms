package max_sub_array

const (
	minInt = ^int(^uint(0) >> 1)
)

// FindMaxSubArray 查找最大连续子数组
// 思路：分治法，
// 首先将规模缩小一半，那么，最大子数组只能存在下面三种情况：
// 1. 位于 [low,mid)子数组中，问题转化为求一个规模较小的相同问题（递归）
// 2. 位于 (mid, high]子数组中，同上
// 3. 跨越两个子数组（包含mid），这种情况，一定是左边从包含mid往左遍历最大的一个子数组 与 mid+1 往右遍历的最大的一个子数组 之和
//    对于第三种情况，时间复杂度是O(n)
// 时间复杂度O(nlogn)
func FindMaxSubArray(a []int, low, high int) (int, int, int) {
	if low < 0 || high > len(a)-1 {
		return -1, -1, 0
	}
	if low == high {
		return low, low, a[low]
	}

	mid := (high + low) / 2

	// 情况1
	leftLow, leftHigh, leftMaxSum := FindMaxSubArray(a, low, mid)
	// 情况2
	rightLow, rightHigh, rightMaxSum := FindMaxSubArray(a, mid+1, high)
	// 情况3
	crossLow, crossHigh, crossMaxSum := findCrossMaxSubArray(a, low, mid, high)

	// 判断是哪一种情况
	if leftMaxSum >= rightMaxSum && leftMaxSum >= crossMaxSum {
		return leftLow, leftHigh, leftMaxSum
	}
	if rightMaxSum >= leftMaxSum && rightMaxSum >= crossMaxSum {
		return rightLow, rightHigh, rightMaxSum
	}
	return crossLow, crossHigh, crossMaxSum
}

func findCrossMaxSubArray(a []int, low, mid, high int) (int, int, int) {
	leftLow := 0
	leftSum := minInt
	sum := 0
	for i := mid; i >= low; i-- {
		sum += a[i]
		if sum > leftSum {
			leftSum = sum
			leftLow = i
		}
	}

	sum = 0
	rightHigh := 0
	rightSum := minInt
	for i := mid + 1; i <= high; i++ {
		sum += a[i]
		if sum > rightSum {
			rightSum = sum
			rightHigh = i
		}
	}
	return leftLow, rightHigh, leftSum + rightSum
}
