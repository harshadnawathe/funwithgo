package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	SENTINEL = math.MaxInt32
)

func mergeSortImpl(arr, ma, mb []int) {
	length := len(arr)
	if length < 2 {
		return
	}

	mid := length / 2
	halfA := arr[0:mid]
	halfB := arr[mid:]

	mergeSortImpl(halfA, ma, mb)
	mergeSortImpl(halfB, ma, mb)

	copy(ma, halfA)
	ma[len(halfA)] = SENTINEL

	copy(mb, halfB)
	mb[len(halfB)] = SENTINEL

	a, b := 0, 0
	for i := range arr {
		if mb[b] < ma[a] {
			arr[i] = mb[b]
			b++
		} else {
			arr[i] = ma[a]
			a++
		}
	}
}

func MergeSort(a []int) {
	ma := make([]int, len(a)+2)
	mid := len(ma) / 2
	mergeSortImpl(a, ma[0:mid], ma[mid:])
}

func main() {
	r := rand.New(rand.NewSource(42))
	arr := r.Perm(11)
	fmt.Println(arr)
	MergeSort(arr)
	fmt.Println(arr)
}
