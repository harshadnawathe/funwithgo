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
}

type handler struct {
	workers pool
	done    chan *worker
}

func newHandler(b Balancer) *handler {
	nWorkers := b.NWorkers
	if nWorkers <= 0 {
		nWorkers = NWorkersDefault
	}

	nWorkersQueueSize := b.NWorkerRequestQueueSize
	if nWorkersQueueSize <= 0 {
		nWorkersQueueSize = NWorkerRequestQueueSizeDefault
	}

	h := new(handler)
	h.done = make(chan *worker)

	for i := 0; i < nWorkers; i++ {
		w := &worker{requests: make(chan Request, nWorkersQueueSize)}
		heap.Push(&h.workers, w)
		w.startWork(h.done)
	}

	return h
}

func (h *handler) dispatch(r Request) {
	w, _ := heap.Pop(&h.workers).(*worker)
	w.requests <- r
	w.pending++
	heap.Push(&h.workers, w)
}

func (h *handler) completed(w *worker) {
	heap.Remove(&h.workers, w.index)
	w.pending--
	heap.Push(&h.workers, w)
}

func (h *handler) stop() {
	for _, w := range h.workers {
		close(w.requests)
	}
	h.workers = nil
}

//Handle starts balacing the load coming from requests channel
func (b Balancer) Handle(requests <-chan Request) {
	h := newHandler(b)
	go func() {
		defer h.stop()
		for {
			select {
			case r, ok := <-requests:
				if !ok {
					return
				}
				h.dispatch(r)
			case w := <-h.done:
				h.completed(w)
			}
		}
	}()
}
