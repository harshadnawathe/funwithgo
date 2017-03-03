package loadbalancer

//worker defines a worker
type worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *worker) startWork(done chan *worker) {
	go func() {
		for r := range w.requests {
			r.C <- r.Fn()
			done <- w
		}
	}()
}
