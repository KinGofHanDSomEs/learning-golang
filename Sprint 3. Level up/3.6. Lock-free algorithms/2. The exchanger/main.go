package main

import (
	"sync/atomic"
)

func AtomicSwap(arg1, arg2 *int32) {
	*arg1, *arg2 = atomic.LoadInt32(arg2), atomic.LoadInt32(arg1)
}

func main() {
	var n1 int32 = 1
	var n2 int32 = 2
	AtomicSwap(&n1, &n2)
	print(n1, " ", n2)
}
