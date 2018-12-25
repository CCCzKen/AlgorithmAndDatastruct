package linkedlist

import (
	"testing"
	"fmt"
)


var list LinkedList


func TestAppend(t *testing.T) {
	if !list.isEmpty() {
		t.Errorf("Linked list should be empty")
	}
	list.Append("first")
	if list.isEmpty() {
		t.Errorf("Linked list should not be empty")
	}
	if size := list.Size(); size != 1 {
		t.Errorf("Wrong count, expected 1 but got %d", size)
	}
	list.Append("second")
	list.Append("third")

	if size := list.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}
}


func TestRemoveAt(t *testing.T) {
	_, err := list.RemoveAt(1)
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
	if size := list.Size(); size != 2 {
		t.Errorf("Wrong count, expected 2 but got %d", size)
	}
}


func TestInsert(t *testing.T) {
	err := list.Insert(1, "second2")
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
	if size := list.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}

	err = list.Insert(0, "zero")
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
}


func TestIndexOf(t *testing.T) {
	if index := list.IndexOf("zero"); index != 0 {
		t.Errorf("Excepted postion 0 but got: %d", index)
	}
	if index := list.IndexOf("first"); index != 1 {
		t.Errorf("Excepted postion 1 but got: %d", index)
	}
	if index := list.IndexOf("second2"); index != 2 {
		t.Errorf("Excepted postion 2 but got: %d", index)
	}
	if index := list.IndexOf("third"); index != 3 {
		t.Errorf("Excepted postion 3 but got: %d", index)
	}
}


func TestHead(t *testing.T) {
	head := list.head
	content := fmt.Sprint(head.content)

	if content != "zero" {
		t.Errorf("Excepted `zero` but got: %s", content)
	}
}