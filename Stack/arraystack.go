package stack

import (
	"container/list"
	"fmt"
)

type ArrayStack struct {
	list *list.List
}

func InitStack() *ArrayStack {
	_list := list.New()
	return &ArrayStack{_list}
}


func (stack *ArrayStack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *ArrayStack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// 获取栈顶
func (stack *ArrayStack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *ArrayStack) Len() int {
	return stack.list.Len()
}


func (stack *ArrayStack) IsEmpty() bool {
	return stack.list.Len() == 0
}

func main() {
	stack := InitStack()
	fmt.Println(stack.IsEmpty())
	stack.Push("czk")
	fmt.Println(stack.Peak())
	fmt.Println(stack.Len())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Len())
	fmt.Println(stack.IsEmpty())
}
