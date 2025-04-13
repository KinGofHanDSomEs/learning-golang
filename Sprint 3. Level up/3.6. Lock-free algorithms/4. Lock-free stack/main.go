package main

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value int
	next  *Node
}

type LockFreeStack struct {
	top unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (l *LockFreeStack) Push(value int) {
	node := &Node{value: value}
	for {
		top := atomic.LoadPointer(&l.top)
		node.next = (*Node)(top)
		if atomic.CompareAndSwapPointer(&l.top, top, unsafe.Pointer(node)) {
			break
		}
	}
}

func (l *LockFreeStack) Pop() (int, bool) {
	for {
		top := atomic.LoadPointer(&l.top)
		if top == nil {
			return 0, false
		}
		node := (*Node)(top)
		if atomic.CompareAndSwapPointer(&l.top, top, unsafe.Pointer(node.next)) {
			return node.value, true
		}
	}
}

func main() {
	stack := NewLockFreeStack()
	for i := 0; i < 3; i++ {
		stack.Push(i)
	}
	for i := 0; i < 3; i++ {
		if val, ok := stack.Pop(); ok {
			println(val)
		} else {
			break
		}
	}
}
