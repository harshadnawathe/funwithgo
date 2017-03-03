package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/harshadnawathe/funwithgo/concurrency/loadbalancer"
)

func main() {
	work := make(chan loadbalancer.Request)
	defer close(work)

	b := loadbalancer.Balancer{NWorkers: 3}
	b.Handle(work)

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out := make(chan int)

			work <- loadbalancer.Request{
				Fn: func() int {
					n := rand.Intn(2000)
					time.Sleep(time.Duration(n) * time.Millisecond)
					return n
				},
				C: out,
			}

			fmt.Println(<-out)
		}()
	}

	wg.Wait()
}
