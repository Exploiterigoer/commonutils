package queueutil

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue(8)
	t.Log(queue.String())
	queue.Put(map[int]string{2: "hello"})
	ok, quantity := queue.Put(map[int]string{1: "welcome to test"})
	t.Log(ok, quantity)
	val, ok, quantity := queue.Get()
	t.Log(val, ok, quantity)
}
