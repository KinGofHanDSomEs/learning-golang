package main  

import (  
	"fmt"  
	"sync/atomic"  
	"unsafe"  
)  

type Node struct {  
	value int  
	next  *Node  
}  

type LockFreeQueue struct {  
	head unsafe.Pointer // Pointer to the first node  
	tail unsafe.Pointer // Pointer to the last node  
}  

// NewLockFreeQueue creates a new LockFreeQueue  
func NewLockFreeQueue() *LockFreeQueue {  
	node := &Node{} // Dummy node  
	queue := &LockFreeQueue{  
		head: unsafe.Pointer(node),  
		tail: unsafe.Pointer(node),  
	}  
	return queue  
}  

// Enqueue adds a value to the queue  
func (q *LockFreeQueue) Enqueue(value int) {  
	newNode := &Node{value: value}  
	for {  
		tail := (*Node)(atomic.LoadPointer(&q.tail))  
		next := tail.next

		if tail == (*Node)(atomic.LoadPointer(&q.tail)) { // Check for tail consistency  
			if next == nil { // If next is nil, try to link new node  
				if atomic.CompareAndSwapPointer(&next, nil, unsafe.Pointer(newNode)) {  
					atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(newNode))  
					return  
				}  
			} else {  
				// Move tail pointer forward if next is not nil  
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(next))  
			}  
		}  
	}  
}  

// Dequeue removes a value from the queue  
func (q *LockFreeQueue) Dequeue() (int, bool) {  
	for {  
		head := (*Node)(atomic.LoadPointer(&q.head))  
		tail := (*Node)(atomic.LoadPointer(&q.tail))  
		next := (*Node)(atomic.LoadPointer(&head.next))  

		if head == (*Node)(atomic.LoadPointer(&q.head)) { // Check for head consistency  
			if head == tail {  
				if next == nil {  
					return 0, false // Queue is empty  
				}  
				// Move tail pointer forward  
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(next))  
			} else {  
				// We have a valid node to dequeue  
				value := next.value  
				if atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(head), unsafe.Pointer(next)) {  
					return value, true  
				}  
			}  
		}  
	}  
}  

func main() {  
	queue := NewLockFreeQueue()  

	// Enqueue some values  
	queue.Enqueue(10)  
	queue.Enqueue(20)  
	queue.Enqueue(30)  

	// Dequeue values  
	for i := 0; i < 3; i++ {  
		val, ok := queue.Dequeue()  
		if ok {  
			fmt.Println(val)  
		} else {  
			fmt.Println("Queue is empty")  
		}  
	}  
}  