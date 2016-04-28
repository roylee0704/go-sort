package main

import "testing"

func TestMerge(t *testing.T) {

	tests := []struct {
		sortedA []int
		sortedB []int
		want    []int
	}{
		{
			sortedA: []int{1},
			sortedB: []int{2},
			want:    []int{1, 2},
		},
		{
			sortedA: []int{1, 8, 9},
			sortedB: []int{2, 4, 5},
			want:    []int{1, 2, 4, 5, 8, 9},
		},
		{
			sortedA: []int{1, 2, 3},
			sortedB: []int{0},
			want:    []int{0, 1, 2, 3},
		},
		{
			sortedA: []int{1, 2, 6, 7},
			sortedB: []int{1, 3, 5, 9},
			want:    []int{1, 1, 2, 3, 5, 6, 7, 9},
		},
	}

	for _, test := range tests {
		mergeSorted := merge(test.sortedA, test.sortedB)
		if !deepCompare(mergeSorted, test.want) {
			t.Errorf("want:%v, got:%v\n", test.want, mergeSorted)
		}
	}

}
