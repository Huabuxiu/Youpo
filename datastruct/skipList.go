package datastruct

type SkipList struct {
	head *SkipListNode

	tail *SkipListNode

	maxLevel int
}

type Level struct {
	nextNode *SkipListNode

	span int64
}

type SkipListNode struct {
	backWard *SkipListNode

	score float64

	level []Level

	value interface{}
}
