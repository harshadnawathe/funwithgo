package loadbalancer

import "container/heap"

type pool []*worker

func (p pool) Len() int {
	return len(p)
}

func (p pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *pool) Swap(i, j int) {
	workers := *p
	workers[i], workers[j] = workers[j], workers[i]
	//note: objects are already swapped - just aasign indices
	workers[i].index = i
	workers[j].index = j
}

func (p *pool) Pop() interface{} {
	workers := *p
	n := len(workers)
	w := workers[n-1]
	*p = workers[0 : n-1]
	return w
}

func (p *pool) Push(w interface{}) {
	wrker, _ := w.(*worker)
	wrker.index = len(*p)
	*p = append(*p, wrker)
}

//NWorkersDefault defines default number of workers created by Balancer
const NWorkersDefault = 8

//NWorkerRequestQueueSizeDefault defines default size of worker queue
const NWorkerRequestQueueSizeDefault = 64

//Balancer defines a Load Balancer
type Balancer struct {
	//NWorkers sets number of workers, if not set then default value will be used
	NWorkers int
	//NWorkerRequestQueueSize sets queue size for a worker, if not set then default value will be used.
	NWorkerRequestQueueSize int
	workers                 pool
	done                    chan *worker
}

func (b *Balancer) initialize() {
	if b.NWorkers == 0 {
		b.NWorkers = NWorkersDefault
	}

	if b.NWorkerRequestQueueSize == 0 {
		b.NWorkerRequestQueueSize = NWorkerRequestQueueSizeDefault
	}

	b.done = make(chan *worker)
}

func (b *Balancer) dispatch(r Request) {
	w, _ := heap.Pop(&b.workers).(*worker)
	w.requests <- r
	w.pending++
	heap.Push(&b.workers, w)
}

func (b *Balancer) completed(w *worker) {
	heap.Remove(&b.workers, w.index)
	w.pending--
	heap.Push(&b.workers, w)
}

func (b *Balancer) stop() {
	for _, w := range b.workers {
		close(w.requests)
	}
	b.workers = nil
}

//Handle starts balacing the load coming from requests channel
func (b *Balancer) Handle(requests <-chan Request) {
	b.initialize()

	for i := 0; i < b.NWorkers; i++ {
		w := &worker{requests: make(chan Request, b.NWorkerRequestQueueSize)}
		heap.Push(&b.workers, w)
		w.startWork(b.done)
	}

	go func() {
		defer b.stop()
		for {
			select {
			case r, ok := <-requests:
				if !ok {
					return
				}
				b.dispatch(r)
			case w := <-b.done:
				b.completed(w)
			}
		}
	}()
}
