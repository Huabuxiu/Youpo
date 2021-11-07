package datastruct

import (
	"fmt"
	"strconv"
	"testing"
)

func printSkipList(skipList *SkipList) {
	skipList.ForEach(func(node interface{}) interface{} {
		print("score : ", fmt.Sprint(node.(*SkipListNode).data.score), " value :", fmt.Sprint(node.(*SkipListNode).data.value))

		var lStr string
		lStr = "[  "
		for i := 0; i < len(node.(*SkipListNode).level) && node != nil; {
			lStr += fmt.Sprint(node.(*SkipListNode).level[i].span, ", ")
			i++
		}

		lStr += "  ]"
		println(lStr)
		return node
	})
	println("maxLevel", skipList.maxLevel, "length :", skipList.Length())
}

func TestSkipList_Insert(t *testing.T) {
	skipList := MakeSkipList()

	for i := 0.0; i < 100.0; i += 1.0 {
		skipList.Insert(i, strconv.FormatFloat(i/10.0, 'f', 1, 64))
	}

	rank := skipList.GetRank(6.0, "0.6")
	println(rank)

	printSkipList(skipList)
}

func TestSkipList_Remove(t *testing.T) {
	skipList := MakeSkipList()

	for i := 0.0; i < 100.0; i += 1.0 {
		skipList.Insert(i, strconv.FormatFloat(i/10.0, 'f', 1, 64))
	}
	printSkipList(skipList)

	skipList.Remove(87.0, "8.7")

	printSkipList(skipList)
}
