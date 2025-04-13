package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// Node представляет узел в стеке.
type Node struct {
	value int           // Значение узла
	next  unsafe.Pointer // Указатель на следующий узел в стеке
}

// Stack представляет lock-free стек.
type Stack struct {
	top unsafe.Pointer // Указатель на вершину стека
}

// Push добавляет новый элемент на вершину стека.
func (s *Stack) Push(value int) {
	node := &Node{value: value}

	for {
		// Загружаем текущую вершину стека
		oldTop := atomic.LoadPointer(&s.top)
		// Устанавливаем новый узел как следующий для добавляемого узла
		node.next = oldTop

		// Если вершину можно атомарно заменить на новый узел, то прерываем цикл
		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(node)) {
			break
		}
	}
}

// Pop удаляет и возвращает элемент с вершины стека. Если стек пуст, возвращает false.
func (s *Stack) Pop() (int, bool) {
	for {
		// Загружаем текущую вершину стека
		oldTop := atomic.LoadPointer(&s.top)
		// Если стек пуст, возвращаем false
		if oldTop == nil {
			return 0, false
		}

		// Загружаем указатель на следующий узел
		newTop := (*Node)(oldTop).next

		// Выполняем атомарную попытку заменить текущую вершину на следующий узел
		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(newTop)) {
			// Возвращаем значение удалённого узла
			return (*Node)(oldTop).value, true
		}
	}
}

func main() {
	// Создаём новый стек
	stack := &Stack{}

	// Добавляем элементы на вершину стека
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Удаляем элементы с вершины стека и выводим их значения
	for {
		if value, ok := stack.Pop(); ok {
			fmt.Println(value)
		} else {
			break
		}
	}
}