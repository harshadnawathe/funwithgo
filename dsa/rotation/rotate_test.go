package rotation

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

func TestShouldRotateGivenIntArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{4, 5, 6, 7, 8, 9, 1, 2, 3}

	f := Rotate(a, 3)

	if f != 6 {
		t.Errorf("Expected is 6 and actual is %v", f)
	}

	if !sliceEq(a, expected) {
		t.Errorf("Expected is %v and actual is %v", expected, a)
	}

}
