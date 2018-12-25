package queue

import (
	"container/list"
	"fmt"
)

type ArrayQueue struct {
	list *list.List
}


func InitQueue() *ArrayQueue {
	_list := list.New()
	return &ArrayQueue{_list}
}

// 入队
func (queue ArrayQueue) EnQueue(value interface{}) {
	queue.list.PushBack(value)
}

// 出队
func (queue ArrayQueue) DeQueue() interface{} {
	e := queue.list.Front()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

// 获取队头
func (queue ArrayQueue) GetHead() interface{}{
	e := queue.list.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

// 获取队尾
func (queue ArrayQueue) GetTail() interface{} {
	e := queue.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}


func (queue ArrayQueue) Size() int {
	return queue.list.Len()
}

func (queue ArrayQueue) IsEmpty() bool {
	return queue.list.Len() == 0
}


func main() {
	queue := InitQueue()
	fmt.Println(queue.IsEmpty())
	queue.EnQueue("chen")
	fmt.Println(queue.list.Len())
	queue.EnQueue("zhao")
	queue.EnQueue("kang")
	fmt.Println(queue.list.Len())
	fmt.Println(queue.GetHead())
	fmt.Println(queue.GetTail())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.GetHead())
	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
}
