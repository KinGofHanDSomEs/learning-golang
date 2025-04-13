package workerpool

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	tasks chan Worker
	wg    sync.WaitGroup
}

func NewPool(mg int) *Pool {
	p := Pool{
		tasks: make(chan Worker),
	}
	for i := 0; i < mg; i++ {
		p.wg.Add(1)
		go func() {
			for task := range p.tasks {
				task.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) AddTask(w Worker) {
	p.tasks <- w
}

func (p *Pool) StopPool() {
	close(p.tasks)
	p.wg.Wait()
}
