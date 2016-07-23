package main

import (
	"fmt"
	"math/rand"

	"github.com/harshadnawathe/go-lang-fun/dsa/partition"
)

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	p := partition.HoarePartitionWithRandomPivot(a, 0)

	QuickSort(a[0:p])
	QuickSort(a[p:])
}

func main() {
	r := rand.New(rand.NewSource(42))
	a := r.Perm(10)
	fmt.Println(a)
	QuickSort(a)
	fmt.Println(a)
}
