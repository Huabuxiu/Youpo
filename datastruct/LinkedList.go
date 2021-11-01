package datastruct

type LinkedList struct {
	//head Node
	head *LinkedListNode

	//tail Node
	tail *LinkedListNode

	//链表长度
	size int
}

type LinkedListNode struct {

	//prev Node
	prev *LinkedListNode

	//next Node
	next *LinkedListNode

	//节点数据
	value interface{}
}

//Add 在链表尾部插入
func (linkedList *LinkedList) Add(data interface{}) {
	linkedList.checkList()

	newNode := &LinkedListNode{
		value: data,
	}

	if linkedList.tail == nil && linkedList.head == nil {
		linkedList.head = newNode
		linkedList.tail = newNode
	} else {
		linkedList.tail.next = newNode
		newNode.prev = linkedList.tail

		linkedList.tail = newNode
	}
	linkedList.size++
}

//Find 查找某一个位置的
func (linkedList *LinkedList) Find(index int) (node *LinkedListNode) {
	linkedList.checkListWithSize(index)

	i := 1
	currentNode := linkedList.head

	for true {
		if i < index {
			currentNode = currentNode.next
			i++
		} else {
			break
		}
	}
	return currentNode
}

//listAddNodeHead 在链表头部插入
func (linkedList *LinkedList) AddAtHead(data interface{}) {
	linkedList.checkList()

	//空链表直接插入
	if linkedList.head == nil {
		linkedList.Add(data)
		return
	}

	//当前节点的下一个节点为head
	newNode := &LinkedListNode{
		value: data,
		next:  linkedList.head,
	}
	linkedList.head = newNode
	linkedList.size++
}

//Set 修改指定位置的值
func (linkedList *LinkedList) Set(index int, data interface{}) {
	node := linkedList.Find(index)
	node.value = data
}

//Insert 在指定位置
func (linkedList *LinkedList) Insert(index int, data interface{}) {
	linkedList.checkListWithSize(index)

	if index == 0 {
		linkedList.AddAtHead(data)
	} else {
		preListNode := linkedList.Find(index - 1)

		newNode := &LinkedListNode{
			value: data,
			next:  preListNode.next,
		}
		preListNode.next.prev = newNode
		preListNode.next = newNode
	}
}

//Remove 指定位置
func (linkedList *LinkedList) Remove(index int) {
	linkedList.checkListWithSize(index)

	if index == 1 {
		linkedList.head = linkedList.head.next
		linkedList.head.prev = nil
	} else if index == linkedList.size {
		linkedList.tail = linkedList.tail.prev
		linkedList.tail.next = nil
	} else {
		listNode := linkedList.Find(index)
		listNode.next.prev = listNode.prev
		listNode.prev.next = listNode.next
	}
}

//RemoveTail
func (linkedList *LinkedList) RemoveTail() {
	linkedList.checkList()
	linkedList.Remove(linkedList.size)
}

//Contains 判断是否存在于链表中
func (linkedList *LinkedList) Contains(value interface{}) bool {
	linkedList.checkList()

	currentNode := linkedList.head

	for i := 1; i < linkedList.size; i++ {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
		i++
	}
	return false
}

//ListSize 链表的长度
func (linkedList *LinkedList) ListSize() (size int) {
	if linkedList == nil {
		panic("list is null")
	}
	return linkedList.size
}

//func (linkedList LinkedList) ForEach(function FunctionInterface) {
//
//}

//make create a new linked list
func Make(values ...interface{}) *LinkedList {
	list := LinkedList{}
	for _, value := range values {
		list.Add(value)
	}
	return &list
}

//check List
func (linkedList *LinkedList) checkList() {
	if linkedList == nil {
		panic("list is null")
	}
}

func (linkedList *LinkedList) checkListWithSize(index int) {
	linkedList.checkList()
	if index > linkedList.size {
		panic("index out of list size")
	}
}
