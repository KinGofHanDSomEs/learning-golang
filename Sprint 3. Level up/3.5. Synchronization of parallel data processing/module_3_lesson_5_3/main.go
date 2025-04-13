package main

import (
	"context"
	"fmt"
	"sync"
)

type Num struct {
	int
}

func (n Num) getSequence() int {
	return n.int
}

type sequenced interface {
	getSequence() int
}

type fanInRecord[T sequenced] struct {
	index int
	data  T
	pause chan struct{}
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	outCh := make(chan T)
	go func() {
		defer close(outCh)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if n.getSequence()%2 == 0 {
					outCh <- n
				}
			}
		}
	}()
	return outCh
}

func OddNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	outCh := make(chan T)
	go func() {
		defer close(outCh)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if n.getSequence()%2 != 0 {
					outCh <- n
				}
			}
		}
	}()
	return outCh
}

func inTemp[T sequenced](
	ctx context.Context,
	channels ...<-chan T,
) <-chan fanInRecord[T] {
	// канал для ожидания
	fanInCh := make(chan fanInRecord[T])
	// для синхронизации
	wg := sync.WaitGroup{}
	// перебор всех входных каналов
	for i := range channels {
		wg.Add(1)
		// запустим горутину для получения данных из канала
		go func(index int) {
			defer wg.Done()
			// канал для синхронизации
			pauseCh := make(chan struct{})
			// цикл для получения данных из канала
			for {
				select {
				// получим данные из канала
				case data, ok := <-channels[index]:
					if !ok {
						return // канал закрыт - выходим
					}
					// положим во временный канал вместе с индексом
					fanInCh <- fanInRecord[T]{
						// индекс канала, откуда пришли данные
						index: index,
						// данные из канала
						data: data,
						// канал для синхронизации
						pause: pauseCh,
					}
				case <-ctx.Done():
					return
				}
				// ждём, пока в канал pause не будет передан сигнал
				// о получении очередного элемента из канала
				select {
				case <-pauseCh:
				// сняли с паузы
				// продолжим обработку данных из входного канала
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		// ожидаем завершения
		wg.Wait()
		close(fanInCh)
	}()
	// вернём канал с неотсортированными элементами
	return fanInCh
}

func processTempCh[T sequenced](
	ctx context.Context,
	inputChannelsNum int,
	fanInCh <-chan fanInRecord[T],
) <-chan T {
	outputCh := make(chan T)
	go func() {
		defer close(outputCh)
		expected := 0
		queuedData := make(map[int]*fanInRecord[T])

		processFromQueue := func() {
			for {
				if item, exists := queuedData[expected]; exists {
					select {
					case outputCh <- item.data:
						item.pause <- struct{}{}
						delete(queuedData, expected)
						expected++
					case <-ctx.Done():
						return
					}
				} else {
					break
				}
			}
		}

		for in := range fanInCh {
			if in.data.getSequence() == expected {
				select {
				case outputCh <- in.data:
					in.pause <- struct{}{}
					expected++
					processFromQueue()
				case <-ctx.Done():
					return
				}
			} else {
				queuedData[in.data.getSequence()] = &in
			}
		}

		processFromQueue()
	}()
	return outputCh
}

func main() {
	nums := []Num{{1}, {2}, {3}, {4}, {5}, {6}}
	fanInCh := inTemp(context.Background(), OddNumbersGen(context.Background(), nums...), EvenNumbersGen(context.Background(), nums...))
	for val := range processTempCh(context.Background(), 2, fanInCh) {
		fmt.Println(val)
	}
}
