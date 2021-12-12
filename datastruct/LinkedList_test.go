package datastruct

import "testing"

func printList(list *LinkedList) {
	println("list size: ", list.size)
	list.ForEach(func(args interface{}) interface{} {
		node := args.(*LinkedListNode)
		println("node value: ", node.value.(int))
		return node
	})
}

func makeTestList() (list *LinkedList) {
	lists := MakeList()
	for i := 0; i < 10; i++ {
		lists.Add(i)
	}
	return lists
}

func TestLinkedList_Add(t *testing.T) {
	list := MakeList()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	printList(list)
}

func TestLinkedList_Find(t *testing.T) {
	list := MakeList()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	printList(list)
	println("--------------")
	for i := 0; i < 10; i++ {
		println("node value: ", list.Find(i).value.(int))
	}

}

func TestLinkedList_AddAtHead(t *testing.T) {
	list := MakeList()
	for i := 0; i < 10; i++ {
		list.AddAtHead(i)
	}
	printList(list)
}

func TestLinkedList_Set(t *testing.T) {
	list := makeTestList()
	printList(list)
	println("------")
	for i := 0; i < 10; i++ {
		list.Set(i, 20-i)
	}
	printList(list)

	//list.Set(-1,20)
	//list.Set(30, 20)
}

func TestLinkedList_Insert(t *testing.T) {
	list := MakeList()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			list.Add(i)
		}
	}
	printList(list)
	println("------")

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			list.Insert(i, i*i)
		}
	}
	printList(list)

}

func TestLinkedList_ListSize(t *testing.T) {
	list := makeTestList()
	print(list.ListSize())
}

func TestLinkedList_Remove(t *testing.T) {
	list := makeTestList()
	printList(list)

	println("-------")

	list.Remove(0)
	list.RemoveTail()
	list.Remove(3)

	printList(list)
}

func TestLinkedList_Contains(t *testing.T) {
	list := makeTestList()
	printList(list)

	println("-------")

	println("list constains 9", list.Contains(9))
	println("list constains 0", list.Contains(0))
	println("list constains 3", list.Contains(3))
	println("list constains -1", list.Contains(-1))
}

func TestName(t *testing.T) {
	m := make(map[string]int)

	m["1234"] = 789
	m["234"] = 73

	delete(m, "12312312321231")

}
