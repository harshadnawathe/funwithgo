package partition

import "github.com/harshadnawathe/go-lang-fun/dsa/rotation"

type Predicate func(int) bool

func StablePartitionRecursive(a []int, pred Predicate) int {
	length := len(a)

	if length == 0 {
		return 0
	}

	if length == 1 {
		if pred(a[0]) {
			return 1
		} else {
			return 0
		}
	}

	mid := length / 2
	n1 := StablePartition(a[0:mid], pred)
	n2 := StablePartition(a[mid:], pred) + mid

	return n1 + rotation.Rotate(a[n1:n2], mid-n1)
}

func StablePartitionWithAuxiliaryArray(a []int, pred Predicate) int {
	aa := make([]int, len(a))
	good, bad := -1, -1
	for i := range a {
		if pred(a[i]) {
			good++
			a[good] = a[i]
		} else {
			bad++
			aa[bad] = a[i]
		}
	}

	for _, b := range aa[0 : bad+1] {
		good++
		a[good] = b
	}

	return len(a) - bad - 1
}

func StablePartition(a []int, pred Predicate) int {
	return StablePartitionWithAuxiliaryArray(a, pred)
}
