package partition

import "testing"

func sliceEq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestShouldStablePartitionGivenIntArray(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pred := func(i int) bool { return (i % 2) == 0 }
	n := StablePartition(a, pred)

	if n != 5 {
		t.Errorf("Expected is 5 and actual is %v", n)
	}

	expected := []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	if !sliceEq(a, expected) {
		t.Errorf("Expected is %v and actaul is %v", expected, a)
	}

}
