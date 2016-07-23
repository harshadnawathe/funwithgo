package partition

func HoarePartitionWithPivot(a []int, begin, pivot_index int) int {
	pivot := a[pivot_index]
	length := len(a)
	a[begin], a[pivot_index] = a[pivot_index], a[begin]

	i, j := begin, length-1
	for {
		for i < length && a[i] <= pivot {
			i++
		}

		for a[j] > pivot {
			j--
		}

		if i >= j {
			a[0], a[j] = a[j], a[0]
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}

func HoarePartitionWithRandomPivot(a []int, begin int) int {
	pivot_index := randomInRange(begin, len(a))
	return HoarePartitionWithPivot(a, begin, pivot_index)
}
