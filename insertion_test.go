package main

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {

	unsorted := []int{5, 4, 3, 2, 1}
	insertion(unsorted)

	for _, v := range unsorted {
		fmt.Println(v)
	}
}

func TestPairWiseSwap(t *testing.T) {

	tests := []struct {
		datas []int
		key   int // key is the divider for sorted/unsorted.
		want  []int
	}{
		{
			datas: []int{0, 1},
			key:   1,
			want:  []int{0, 1},
		},
		{
			datas: []int{1, 0},
			key:   1,
			want:  []int{0, 1},
		},
		{
			datas: []int{2, 9, 3, 1},
			key:   2,
			want:  []int{2, 3, 9, 1},
		},
	}

	for _, test := range tests {
		pairwiseSwap(test.datas, test.key)

		if !deepCompare(test.want, test.datas) {
			t.Errorf("want: %v, got: %v", test.want, test.datas)
		}
	}

}

func deepCompare(a []int, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
