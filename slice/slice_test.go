package slice

import (
	"testing"
)

var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestPop(t *testing.T) {
	i, altered := Pop(slice)
	if i != 9 {
		t.Errorf("failed Pop with value of %d when 9 was expected", i)
		return
	} else {
		t.Logf("passed Pop popped check with value of %d", i)
	}

	if l := len(altered); l != 9 {
		t.Errorf("failed Pop with length of %d when 9 was expected", l)
		return
	} else {
		t.Logf("passed Pop length check with length of %d", l)
	}

	t.Logf("passed Pop test with slice of %v", altered)
}

func TestMap(t *testing.T) {
	altered := Map(slice, func(_ int, v int) int {
		return v + 1
	})
	t.Logf("Mapped %v to %v", slice, altered)
}

func TestReduce(t *testing.T) {
	altered := Reduce(slice, func(_ int, v int) bool {
		return v%2 == 0
	})
	t.Logf("Reduced %v to %v", slice, altered)
}
