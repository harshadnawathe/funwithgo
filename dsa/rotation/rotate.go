package rotation

func Rotate(a []int, nFirst int) int {
	first, next, last := 0, nFirst, len(a)
	result := first + (last - nFirst)

	for first != next {
		a[first], a[next] = a[next], a[first]

		first++
		next++

		if next == last {
			next = nFirst
		} else if first == nFirst {
			nFirst = next
		}
	}

	return result
}
