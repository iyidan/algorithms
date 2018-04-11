package skiplist

import (
	"math/rand"
	"sync"
)

const (
	// MaxLevel max level of skiplist
	MaxLevel = 32
	// RandMax commonly use 1/4 or 1/2
	RandMax = uint32(float32(1<<32) * 0.25)
)

// Skiplist skiplist implement
// see https://zh.wikipedia.org/wiki/%E8%B7%B3%E8%B7%83%E5%88%97%E8%A1%A8
type Skiplist struct {
	head, tail  *skiplistNode
	length      uint                       // current length
	level       int                        // current max level
	scoreMap    map[interface{}]float64    // O(1) to get node's score
	lock        sync.RWMutex               // insure threadsafe
	compareFunc func(a, b interface{}) int // func for compare data when the elem score are the same
}

// Node node only store score and data
type Node struct {
	Score float64
	Data  interface{} // Skiplist.scoreMap's key
}

type skiplistNode struct {
	backward *skiplistNode
	score    float64
	data     *interface{} // pointer to Skiplist.scoreMap's key
	level    []skiplistLevel
}

type skiplistLevel struct {
	span    uint
	forward *skiplistNode
}

// New skiplist generator
func New(cf func(a, b interface{}) int) *Skiplist {
	sl := Skiplist{compareFunc: cf}
	sl.init()
	return &sl
}

// rand a level
func randomLevel() (level int) {
	level = 1
	for level < MaxLevel && rand.Uint32() < RandMax {
		level++
	}
	return level
}

// init head node
func (s *Skiplist) init() {
	s.scoreMap = make(map[interface{}]float64)

	head := skiplistNode{}
	head.level = make([]skiplistLevel, MaxLevel)
	s.head = &head
	s.level = 1
	s.length = 0
}

// Add if node exists, only update its score
// if add/update success, return true
func (s *Skiplist) Add(data interface{}, score float64) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	isUpdate := false
	dataScore, ok := s.scoreMap[data]
	if ok {
		if dataScore == score {
			return false
		}
		isUpdate = true
	}

	update := [MaxLevel]*skiplistNode{}
	rank := [MaxLevel]uint{}
	if isUpdate {
		s.delByScore(data, dataScore)
	}

	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		if i == s.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && s.compareFunc(*x.level[i].forward.data, data) < 0)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	level := randomLevel()
	x = &skiplistNode{data: &data, score: score}
	x.level = make([]skiplistLevel, level)

	// when rand level gt s.level
	if level > s.level {
		for i := s.level; i < level; i++ {
			rank[i] = 0
			update[i] = s.head
			update[i].level[i].span = s.length
		}
		s.level = level
	}

	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	// level < s.level
	for i := level; i < s.level; i++ {
		update[i].level[i].span++
	}

	// backward
	if update[0] != s.head {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		s.tail = x
	}

	s.length++
	// add to map
	s.scoreMap[data] = score

	return true
}

// if data not exists return false
func (s *Skiplist) delByScore(data interface{}, score float64) bool {
	// search node
	x := s.head
	update := make([]*skiplistNode, s.level)
	for i := s.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && s.compareFunc(*x.level[i].forward.data, data) < 0)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if x != nil && x.score == score && s.compareFunc(*x.data, data) == 0 {
		s.delNode(x, update)
		return true
	}
	return false
}

func (s *Skiplist) delNode(x *skiplistNode, update []*skiplistNode) {
	for i := 0; i < s.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		s.tail = x.backward
	}

	for s.level > 1 && s.head.level[s.level-1].forward == nil {
		s.level--
	}
	s.length--
	delete(s.scoreMap, *x.data)
}

// Del return the real delnum
func (s *Skiplist) Del(datas ...interface{}) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	i := 0
	for _, v := range datas {
		dataScore, ok := s.scoreMap[v]
		if ok {
			if s.delByScore(v, dataScore) {
				i++
			}
		}
	}
	return i
}

// find the left node which score >= give score [score, ...)
func (s *Skiplist) findMinScoreNode(score float64) *skiplistNode {
	if s.length == 0 {
		return nil
	}
	if s.tail.score < score {
		return nil
	}

	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.score < score {
			x = x.level[i].forward
		}
	}

	// x is the letf-nearest node for the give score
	if x.level[0].forward != nil && x.level[0].forward.score >= score {
		return x.level[0].forward
	}
	return nil
}

// RangeByScore get the node's data with scores in [minScore, maxScore)
func (s *Skiplist) RangeByScore(minScore, maxScore float64) (ret []*Node) {
	if minScore >= maxScore {
		return
	}
	s.lock.RLock()
	defer s.lock.RUnlock()

	minNode := s.findMinScoreNode(minScore)
	if minNode == nil {
		return
	}

	ret = make([]*Node, 0)
	for minNode != nil && minNode.score < maxScore {
		ret = append(ret, &Node{Score: minNode.score, Data: *minNode.data})
		minNode = minNode.level[0].forward
	}
	return
}

// Rank get node rank, rank > 0,
// if data not exists, return 0
func (s *Skiplist) Rank(data interface{}) (rank uint) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	// get score
	score, ok := s.scoreMap[data]
	if !ok {
		return 0
	}

	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && s.compareFunc(*x.level[i].forward.data, data) < 0)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}
		if x.level[i].forward != nil && x.level[i].forward.score == score && s.compareFunc(*x.level[i].forward.data, data) == 0 {
			rank += x.level[i].span
			break
		}
	}
	return rank
}

// Score get node score
func (s *Skiplist) Score(data interface{}) (score float64, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	score, ok = s.scoreMap[data]
	return
}

// RangeByRank get the range of nodes in [minRank, maxRank)
func (s *Skiplist) RangeByRank(minRank, maxRank uint) (ret []*Node) {
	if minRank >= maxRank {
		return
	}
	s.lock.RLock()
	defer s.lock.RUnlock()

	x := s.head
	rank := uint(0)
	for i := s.level - 1; i >= 0; i-- {
		rank += x.level[i].span
		for x.level[i].forward != nil && rank < minRank {
			x = x.level[i].forward
		}
	}
	if x.level[0].forward == nil {
		return
	}

	rank = minRank
	ret = make([]*Node, 0)
	for x.level[0].forward != nil && rank < maxRank {
		ret = append(ret, &Node{Score: x.level[0].forward.score, Data: *x.level[0].forward.data})
		rank++
		x = x.level[0].forward
	}
	return
}

// Len return the length of current skiplist
func (s *Skiplist) Len() uint {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.length
}
