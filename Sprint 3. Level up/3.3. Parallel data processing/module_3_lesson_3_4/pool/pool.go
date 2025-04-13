package pool

import (
	"errors"
	"sync"
)

var (
	Err = errors.New("error")
)

type PoolTask interface {
	Execute() error
	OnFailure(error)
}

type WorkerPool interface {
	Start()
	Stop()
	AddWork(PoolTask)
}

type MyPool struct {
	tasks      chan PoolTask
	numWorkers int
	mu         sync.Mutex
	wg         sync.WaitGroup
	startOnce  sync.Once
	stopOnce   sync.Once
	isStarted  bool
	done       chan struct{}
}

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error) {
	if channelSize < 1 || numWorkers < 1 {
		return nil, Err
	}
	return &MyPool{
		numWorkers: numWorkers,
		tasks:      make(chan PoolTask, channelSize),
	}, nil
}

func (mp *MyPool) Start() {
	mp.startOnce.Do(func() {
		mp.mu.Lock()
		mp.isStarted = true
		mp.mu.Unlock()

		for i := 0; i < mp.numWorkers; i++ {
			mp.wg.Add(1)
			go mp.worker()
		}
	})
}

func (mp *MyPool) Stop() {
	mp.startOnce.Do(func() {
		mp.mu.Lock()
		mp.isStarted = false
		mp.mu.Unlock()

		close(mp.done)
		mp.wg.Wait()
		close(mp.tasks)
	})
}

func (mp *MyPool) AddWork(p PoolTask) {
	mp.mu.Lock()
	if !mp.isStarted {
		return
	}
	mp.mu.Unlock()
	select {
	case mp.tasks <- p:
	case <-mp.done:
	}
}

func (mp *MyPool) worker() {
	defer mp.wg.Done()
	for {
		select {
		case task, ok := <-mp.tasks:
			if !ok {
				return
			}
			if err := task.Execute(); err != nil {
				task.OnFailure(err)
			}
		case <-mp.done:
			return
		}
	}
}
