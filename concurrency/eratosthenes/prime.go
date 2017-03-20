package eratosthenes

//PrimeSieve describes a Eratosthenes Prime Sieve
type PrimeSieve interface {
	Chan() <-chan int
	Close()
}

//NewPrimeSieve creates a prime sieve
func NewPrimeSieve() PrimeSieve {
	return defaultSieve()
}

type primeSieve struct {
	primes <-chan int
	done   chan struct{}
}

func (p *primeSieve) Chan() <-chan int {
	return p.primes
}

func (p *primeSieve) Close() {
	close(p.done)
}

//defaultSieve creates an infinite prime number sieve
func defaultSieve() *primeSieve {
	ps := new(primeSieve)
	ps.primes, ps.done = sieve()
	return ps
}

func sieve() (<-chan int, chan struct{}) {
	primes := make(chan int)
	done := make(chan struct{})

	ch := generator(done)
	go func() {
		for {
			if prime, ok := <-ch; ok {
				primes <- prime
				chf := filter(ch, func(n int) bool {
					return n%prime != 0
				})
				ch = chf
			} else {
				break
			}
		}
		close(primes)
	}()
	return primes, done
}

func generator(done chan struct{}) <-chan int {
	nums := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case nums <- i:
			case <-done:
				close(nums)
				return
			}
		}
	}()
	return nums
}

func filter(nums <-chan int, f func(int) bool) <-chan int {
	fnums := make(chan int)
	go func() {
		for n := range nums {
			if f(n) {
				fnums <- n
			}
		}
		close(fnums)
	}()
	return fnums
}
