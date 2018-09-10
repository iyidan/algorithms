package zero_sub_array

import (
	"reflect"
	"testing"
)

func TestFindAllZeroSubArray(t *testing.T) {
	type testCase struct {
		arr []int
		ret [][2]int
	}

	cases := []testCase{
		testCase{
			arr: []int{},
			ret: nil,
		},
		testCase{
			arr: []int{1},
			ret: nil,
		},
		testCase{
			arr: []int{0},
			ret: [][2]int{[2]int{0, 1}},
		},
		testCase{
			arr: []int{1, 2, 3, 4, 5, 6, -11, 7, 0, 9},
			ret: [][2]int{[2]int{4, 7}, [2]int{8, 9}},
		},
		testCase{
			arr: []int{1, 1, 1, 1, 5, 0, -5, 7, 1, -7, -1, 8, 0, 9},
			ret: [][2]int{[2]int{5, 6}, [2]int{4, 7}, [2]int{4, 11}, [2]int{7, 11}, [2]int{9, 12}, [2]int{9, 13}, [2]int{12, 13}},
		},
	}

	for _, c := range cases {
		ret := FindAllZeroSubArray(c.arr)
		t.Logf("arr: %v, ret: %v, correct: %v", c.arr, ret, c.ret)
		if !reflect.DeepEqual(ret, c.ret) {
			t.Fatalf("arr: %v, ret: %v, correct: %v", c.arr, ret, c.ret)
		}
	}
}
