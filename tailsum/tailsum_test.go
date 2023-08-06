package tailsum_test

import (
	"ntail/tailsum"
	"testing"
)

type Test struct {
	in  []int
	sum int
	out int
}

var tests = []Test{
	{[]int{1, 2, -1, 0, 5}, 4, 3},
	{[]int{1, 2, 3, 4, 5}, 12, 3},
	{[]int{1, 2, 3, 4, 5}, 15, 5},
	{[]int{1, 2, 3, 4, 5}, 8, 0},
	{[]int{1, -1, 1, -1, 1}, 0, 2},
	{[]int{}, 23, 0},
	{[]int{3, 0, -2, 4, -6}, -4, 3},
}

func TestTailSum(t *testing.T) {
	for i, test := range tests {
		size := tailsum.TailSum(test.in, test.sum)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%d; want %d", i, test.in, size, test.out)
		}
	}
}
