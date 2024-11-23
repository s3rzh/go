package main

import (
	"fmt"
	"sync"
	"time"
)

// Предположим, нескольким горутинам нужен монопольный доступ к общему ресурсу. 
// sync.Cond можно использовать для координации доступа. 
// Например, пулу рабочих горутин может потребоваться подождать, пока определенное количество ресурсов не станет доступным, прежде чем они смогут начать обработку. 
// Горутины могут ожидать условной переменной, используя cond.Wait() , и уведомлять об освобождении ресурса, используя cond.Signal() или cond.Broadcast() .

const MaxResources = 3

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	resourceProvider := NewResourceProvider(cond, MaxResources)

	wg.Add(10)

	for i := range 10 {
		go func(workerID int) {
			defer wg.Done()
			worker := NewWorker(workerID, cond, resourceProvider)
			worker.Run()
		}(i)
	}

	wg.Wait()
}

type ResourceProvider struct {
	maxResources       int
	availableResources int
	cond               *sync.Cond
}

func NewResourceProvider(cond *sync.Cond, maxResources int) *ResourceProvider {
	return &ResourceProvider{
		cond:               cond,
		availableResources: maxResources,
	}
}

func (rp *ResourceProvider) AvailableResources() int {
	return rp.availableResources
}

func (rp *ResourceProvider) AcquireResoirce() {
	rp.availableResources--
}

func (rp *ResourceProvider) ReleaseResource() {
	rp.availableResources++
}

type Worker struct {
	id   int
	cond *sync.Cond
	rp   *ResourceProvider
}

func NewWorker(workerID int, cond *sync.Cond, rp *ResourceProvider) *Worker {
	return &Worker{
		id:   workerID,
		cond: cond,
		rp:   rp,
	}
}

func (w *Worker) Run() {
	w.cond.L.Lock()
	for w.rp.AvailableResources() == 0 {
		fmt.Printf("Worker %d is waiting for resources\n", w.id)
		w.cond.Wait()
	}

	w.rp.AcquireResoirce() // Взятие ресурса

	fmt.Printf("Worker %d acquired resource. Remaining resources: %d\n", w.id, w.rp.AvailableResources())
	w.cond.L.Unlock()
	time.Sleep(1 * time.Second) // Simulating work
	w.cond.L.Lock()
	defer w.cond.L.Unlock()

	w.rp.ReleaseResource() // Освобождение ресурса

	fmt.Printf("Worker %d released resource. Remaining resources: %d\n", w.id, w.rp.AvailableResources())
	w.cond.Signal()
}

// В этом примере у нас есть несколько рабочих горутин, которым требуется монопольный доступ к ограниченным ресурсам. Рабочие горутины приобретают и освобождают ресурсы, используя cond.Signal() для координации с другими рабочими процессами. 
// Если ресурсы недоступны, рабочие горутины using cond.Wait() ждут, пока другая горутина не освободит ресурс.

// В этом примере sync.Cond обеспечивает синхронизацию и координацию между рабочими горутинами, гарантируя, что рабочие горутины будут ждать, когда ресурсы недоступны, тем самым эффективно синхронизируя доступ к ресурсам.
