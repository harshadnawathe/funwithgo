package maxSum

import "math"

func max(a int, as ...int) int {
	if len(as) == 0 {
		return a
	}

	result := a
	for _, each := range as {
		if result < each {
			result = each
		}
	}
	return result
}

func max2(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func MaxSumQuad(a []int) int {
	n := len(a)

	if n == 0 {
		return 0
	}

	maxSum := math.MinInt32

	for i := range a {
		sum := 0
		for j := i; j < n; j++ {
			sum += a[j]
			//maxSum = max2(maxSum, sum)
			if maxSum < sum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func MaxSumLinear(a []int) int {
	if len(a) == 0 {
		return 0
	}

	maxSumSoFar, maxEndingHere := 0, 0
	for _, e := range a {
		maxEndingHere = max2(e+maxEndingHere, math.MinInt32)
		maxSumSoFar = max2(maxSumSoFar, maxEndingHere)
	}
	return maxSumSoFar
}

func MaxSumDivideAndConquer(a []int) int {
	length := len(a)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return a[0]
	}

	m := length / 2

	maxLeft := math.MinInt32
	sumLeft := 0
	for i := m - 1; i >= 0; i-- {
		sumLeft += a[i]
		maxLeft = max2(maxLeft, sumLeft)
	}

	maxRight := math.MinInt32
	sumRight := 0
	for i := m; i < length; i++ {
		sumRight += a[i]
		maxRight = max2(maxRight, sumRight)
	}
	return max(maxLeft+maxRight, MaxSumDivideAndConquer(a[0:m]), MaxSumDivideAndConquer(a[m:]))
}
