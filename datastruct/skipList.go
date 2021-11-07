package datastruct

import (
	"github.com/Huabuxiu/Youpo/define"
	"math/rand"
)

const (
	MaxLevelSize = 32
	p            = 0.25
)

type SkipList struct {
	head *SkipListNode

	tail *SkipListNode

	maxLevel int

	length int64
}

type Level struct {
	nextNode *SkipListNode

	//跨度，表示走到下一个节点，需要经过几个节点
	//从初始节点到当前节点的跨度直接定义为 当前节点的排位
	span int64
}

type Item struct {

	//存储的值
	value interface{}

	//排序分值
	score float64
}

type SkipListNode struct {
	data Item

	backWard *SkipListNode

	level []Level
}

func (skipList *SkipList) randomLevel() int {
	result := 1
	for rand.Float64() < p && result < MaxLevelSize {
		result += 1
	}
	return result
}

func (skipList SkipList) check() {
	if skipList.head == nil {
		panic("skip list not init")
	}
}

func MakeSkipList() *SkipList {
	return &SkipList{
		maxLevel: 1,
		length:   0,
		head:     MakeSkipListNode(MaxLevelSize, nil, 0),
	}
}

func MakeSkipListNode(level int, value interface{}, score float64) *SkipListNode {
	skipListNode := SkipListNode{
		level: make([]Level, level),
		data: Item{
			value: value,
			score: score,
		},
	}

	for i := range skipListNode.level {
		skipListNode.level[i] = Level{
			span: 0,
		}
	}

	return &skipListNode
}

func (skipList *SkipList) Length() int64 {
	return skipList.length
}

func (skipList *SkipList) ForEach(function define.FunctionInterface) {
	curNode := skipList.head

	for curNode != nil {
		function(curNode)
		curNode = curNode.level[0].nextNode
	}

}

func (skipList *SkipList) Insert(score float64, value interface{}) {
	skipList.check()

	//生成当前层数
	randomLevel := skipList.randomLevel()

	if randomLevel > skipList.maxLevel {
		skipList.maxLevel = randomLevel
	}

	//初始化node
	node := MakeSkipListNode(randomLevel, value, score)

	preNode := make([]*SkipListNode, randomLevel)
	preNodeRank := make([]int64, randomLevel)

	//获取当前节点每层level 的前一个节点
	for i := 0; i < randomLevel; i++ {
		currNode := skipList.head

		for currNode.level[i].nextNode != nil &&
			currNode.level[i].nextNode.data.score < score {
			preNodeRank[i] += currNode.level[i].span
			currNode = currNode.level[i].nextNode
		}
		preNode[i] = currNode
	}

	for i, preNodeItem := range preNode {
		//前一个节点的next 为 nil
		if preNodeItem.level[i].nextNode == nil {
			preNodeItem.level[i].nextNode = node
			//当前层的前驱节点到当前层的距离 等于 第0层前驱节点的值+1 - 当前层前驱节点的Rank
			preToCurSpan := preNodeRank[0] + 1 - preNodeRank[i]

			//当前节点到nil 的span 是原来 前置节点到 next 的span -  原来 前置节点到 当前节点的span
			node.level[i].span = 0
			preNodeItem.level[i].span = preToCurSpan
		} else {
			//链表连接
			node.level[i].nextNode = preNodeItem.level[i].nextNode
			preNodeItem.level[i].nextNode = node

			preToNextSpan := preNodeItem.level[i].span
			preToCurSpan := preNodeRank[0] + 1 - preNodeRank[i]

			//当前节点到nil 的span 是原来 前置节点到 next 的span -  原来 前置节点到 当前节点的span
			node.level[i].span = preToNextSpan - preToCurSpan
			preNodeItem.level[i].span = preToCurSpan
		}
	}
	node.backWard = preNode[0]

	if skipList.tail == nil || node.backWard == skipList.tail {
		skipList.tail = node
	}

	skipList.length++

}

func (skipList *SkipList) Remove(score float64, value interface{}) {
	skipList.check()

	node := skipList.getNode(score, value)
	if node == nil {
		return
	}

	//找到每层的前驱节点
	preNode := make([]*SkipListNode, len(node.level))

	skipList.removeNode(node, preNode)
}

func (skipList *SkipList) removeNode(node *SkipListNode, preNode []*SkipListNode) {
	for i := len(node.level) - 1; i >= 0; i-- {
		currNode := skipList.head

		for currNode.level[i].nextNode.data != node.data {
			currNode = currNode.level[i].nextNode
		}

		preNode[i] = currNode
	}

	for i := len(node.level) - 1; i >= 0; i-- {
		preNode[i].level[i].nextNode = node.level[i].nextNode
		preNode[i].level[i].span += node.level[i].span - 1
	}

	//尾节点
	if node.level[0].nextNode == nil {
		skipList.tail = node.backWard
	} else {
		node.level[0].nextNode.backWard = preNode[0]
	}

	//如果删除了最高层的节点
	if len(node.level) == skipList.maxLevel {
		for skipList.head.level[skipList.maxLevel].nextNode == nil {
			skipList.maxLevel--
		}
	}

	skipList.length--
}

func (skipList *SkipList) GetRank(score float64, value interface{}) int64 {
	skipList.check()
	rank, node := skipList.findNodeAndGetRank(score, value)
	if node != nil && rank != 0 {
		return rank
	}

	return 0
}

func (skipList *SkipList) findNodeAndGetRank(score float64, value interface{}) (int64, *SkipListNode) {
	var rank int64 = 0

	currNode := skipList.head
	for i := skipList.maxLevel - 1; i >= 0; i-- {

		for currNode.level[i].nextNode != nil &&
			(currNode.level[i].nextNode.data.score < score ||
				//只有score 和 value 相等才会往前移动
				currNode.level[i].nextNode.data.score == score && value == currNode.level[i].nextNode.data.value) {
			// 只移动了currNode 层数还没有换
			rank = currNode.level[i].span + rank
			currNode = currNode.level[i].nextNode
		}

		if currNode.data.value == value {
			return rank, currNode
		}

	}
	return 0, nil
}

func (skipList *SkipList) getNode(score float64, value interface{}) *SkipListNode {
	skipList.check()
	rank, node := skipList.findNodeAndGetRank(score, value)
	if node != nil && rank != 0 {
		return node
	}
	return nil
}

func (skipList *SkipList) GetByRank(rank int64) *SkipListNode {
	return nil
}

func (skipList *SkipList) HasInRange(start float64, end float64) bool {
	return false
}

func (skipList *SkipList) GetFirstInRange(start float64, end float64) (float64, interface{}) {
	return 0, 1
}

func (skipList *SkipList) GetLastInRange(start float64, end float64) (float64, interface{}) {
	return 0, 1
}

func (skipList *SkipList) RemoveInRange(start float64, end float64) {

}

func (skipList *SkipList) RemoveInRangeByRank(start int64, end int64) {

}
