package loadbalancer

//Request defines a Job request to be executed
type Request struct {
	Fn func() int
	C  chan int
}
