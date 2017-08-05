package dst

// BTSortDataer 可排序的数据定义
type BTSortDataer interface {
	Compare(data BTSortDataer) int
}

// BTDStr string alias
type BTDStr string

// Compare implement BTSortDataer
func (btd BTDStr) Compare(b BTSortDataer) int {
	if btd == b {
		return 0
	}
	if btd > b.(BTDStr) {
		return 1
	}
	return -1
}

// BTDInt int alias
type BTDInt int

// Compare implement BTSortDataer
func (btd BTDInt) Compare(b BTSortDataer) int {
	if btd == b {
		return 0
	}
	if btd > b.(BTDInt) {
		return 1
	}
	return -1
}
