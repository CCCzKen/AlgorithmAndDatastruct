package queue

import "testing"

func TestArrayStack(t *testing.T) {
	queue := InitQueue()
	queue.EnQueue("chen")
	queue.EnQueue("zhao")
	queue.EnQueue("kang")
	queue.EnQueue(1)

	len := queue.Size()
	if len != 4 {
		t.Errorf("queue.Size() failed. Got %d, expected 4.", len)
	}

	value := queue.GetHead()
	if value != "chen" {
		t.Errorf("queue.GetHead() failed. Got %s, expected chen.", value)
	}

	value = queue.GetTail()
	if value != 1 {
		t.Errorf("queue.GetTail() failed. Got %d, expected 1.", value)
	}

	value = queue.DeQueue()
	if value != "chen" {
		t.Errorf("queue.DeQueue() failed. Got %s, expected chen.", value)
	}

	len = queue.Size()
	if len != 3 {
		t.Errorf("queue.Size() failed. Got %d, expected 3.", len)
	}

	value = queue.GetHead()
	if value != "zhao" {
		t.Errorf("queue.GetHead() failed. Got %s, expected zhao.", value)
	}

	value = queue.DeQueue()
	if value != "zhao" {
		t.Errorf("queue.DeQueue() failed. Got %s, expected zhao.", value)
	}

	value = queue.DeQueue()
	if value != "kang" {
		t.Errorf("queue.DeQueue() failed. Got %s, expected kang.", value)
	}

	empty := queue.IsEmpty()
	if empty {
		t.Errorf("queue.IsEmpty() failed.")
	}

	value = queue.DeQueue()
	if value != 1 {
		t.Errorf("queue.DeQueue() failed. Got %d, expected 1.", value)
	}

	empty = queue.IsEmpty()
	if !empty {
		t.Errorf("queue.IsEmpty() failed.")
	}

	len = queue.Size()
	if len != 0 {
		t.Errorf(" ueue.Size() failed. Got %d, expected 0.", len)
	}
}