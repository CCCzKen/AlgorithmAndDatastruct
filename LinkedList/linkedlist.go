package linkedlist

import (
	"sync"
	"fmt"
)

type Item interface {}

type Node struct {
	content Item
	next *Node
}


type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}


// 在链表结尾追加元素
func (list *LinkedList) Append(text string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	newNode := Node{text, nil}

	if list.head == nil {
		list.head = &newNode
	} else {
		curNode := list.head
		for {
			if curNode.next == nil {
				break
			}
			curNode = curNode.next
		}
		curNode.next = &newNode
	}

	list.size++
}


// 在链表指定位置插入指定元素
func (list *LinkedList) Insert(index int, text string) error {
	list.lock.Lock()
	defer list.lock.Unlock()

	if index < 0 || index > list.size {
		return fmt.Errorf("Index %d out of bonuds", index)
	}
	newNode := Node{text, nil}

	// 插入链表头部
	if index == 0 {
		newNode.next = list.head
		list.head = &newNode
		list.size++
		return nil
	}

	preNode := list.head
	preIndex := 0
	for preIndex <= index - 2 {
		preIndex++
		preNode = preNode.next
	}
	newNode.next = preNode.next
	preNode.next = &newNode
	list.size++
	return nil
}


// 删除指定位置的元素
func (list *LinkedList) RemoveAt(index int) (*Item, error){
	list.lock.Lock()
	defer list.lock.Unlock()

	if index < 0 || index > list.size {
		return nil, fmt.Errorf("Index %d out of bonuds", index)
	}

	curNode := list.head
	preIndex := 0
	for preIndex < index - 1 {
		preIndex++
		curNode = curNode.next
	}

	item := curNode.content
	curNode.next = curNode.next.next
	list.size--
	return &item, nil
}


// // 获取指定元素在链表中的索引
func (list *LinkedList) IndexOf(item Item) int {
	list.lock.Lock()
	defer list.lock.Unlock()
	curNode := list.head
	localIndex := 0
	for {
		if curNode.content == item {
			return localIndex
		}
		if curNode.next == nil {
			return -1
		}
		curNode = curNode.next
		localIndex++
	}
}


// 检查链表是否为空
func (list *LinkedList) isEmpty() bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.head == nil {
		return true
	}
	return false
}


// 获取链表的长度
func (list *LinkedList) Size() int {
	list.lock.Lock()
	defer list.lock.Unlock()
	size := 1
	nextNode := list.head
	for {
		if nextNode == nil || nextNode.next == nil {
			break
		}
		size++
		nextNode = nextNode.next
	}
	return size
}


// 输出所有链表
func (list *LinkedList) ShowAll() {
	if list.Size() == 1 {
		fmt.Println(list.head.content)
	} else {
		curNode := list.head
		for {
			fmt.Println(curNode.content)
			if curNode.next == nil {
				break
			} else {
				curNode = curNode.next
			}
		}
	}
}


func main() {
	link := LinkedList{nil, 0, sync.RWMutex{}}
	fmt.Println(link.isEmpty())
	link.Append("chen")
	link.Append("zhao")
	link.Append("kang")
	link.Insert(2, "xxx")
	fmt.Println(link.Size())
	fmt.Println(link.isEmpty())

	link.RemoveAt(2)
	link.ShowAll()

	fmt.Println(link.IndexOf("chen"))
}
