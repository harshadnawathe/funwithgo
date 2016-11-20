package maxSum

import "testing"

func assert(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected max sum of subsequence to be %v, found %v", expected, actual)
	}
}
func TestShouldReturnMaxSumOfSubsequence(t *testing.T) {

	functionUnderTest := MaxSumDivideAndConquer

	assert(t, 0, functionUnderTest([]int{}))

	assert(t, 6, functionUnderTest([]int{6}))

	assert(t, 1, functionUnderTest([]int{1, 0}))
	assert(t, 1, functionUnderTest([]int{-1, 1}))
	assert(t, 2, functionUnderTest([]int{1, 1}))

	assert(t, 4, functionUnderTest([]int{4, -1, -2}))
	assert(t, 2, functionUnderTest([]int{-1, 2, 0}))
	assert(t, 5, functionUnderTest([]int{-1, -1, 5}))
	assert(t, 3, functionUnderTest([]int{2, 1, -2}))
	assert(t, 6, functionUnderTest([]int{-2, 4, 2}))
	assert(t, 12, functionUnderTest([]int{2, 4, 6}))

	assert(t, 187, functionUnderTest([]int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}))
}

func TestShouldHandleAllNegativeNumberSequence(t *testing.T) {
	//assert(t, -1, MaxSum([]int{-2, -1, -3}))
	assert(t, -1, MaxSumDivideAndConquer([]int{-2, -1, -3}))
	assert(t, -1, MaxSumQuad([]int{-2, -1, -3}))
}
