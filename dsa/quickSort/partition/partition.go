package partition

import (
	"math/rand"
	"time"
)

var randGen = rand.New(rand.Seed(time.Now().Unix()))

func randomInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func LomutoPartitionWithPivot(a []int, begin, pivot_index int) int {
	pivot := a[pivot_index]

	last := len(a) - 1
	a[pivot_index], a[last] = a[last], a[pivot_index]

	i, j := begin, begin
	for ; i < last; i++ {
		if a[i] <= pivot {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}
	a[j], a[last] = a[last], a[j]
	return j
}

func LomutoPartitionWithRandomPivot(a []int, begin int) int {
	pivot_index = randomInRange(begin, len(a))
	return LomutoPartitionWithPivot(a, begin, pivot_index)
}
