package max_sub_array

import (
	"testing"
)

func TestFindMaxSubArray(t *testing.T) {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, sum := FindMaxSubArray(a, 0, len(a)-1)
	if low != 7 || high != 10 || sum != 43 {
		t.Fatal(low, high, sum)
	}
}
