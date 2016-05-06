package main

import "fmt"

func main() {
	fmt.Println(divide([]int{10, 8, 3}))

}
func divide(input []int) []int {

	if len(input) == 1 {
		return input
	}

	a := divide(input[:len(input)/2])
	b := divide(input[len(input)/2:])
	return merge(a, b)
}

// merge uses 2 fingers algorithm to merge 2 sorted arrays.
// O(|a| + |b|) = O(n) time complexity
func merge(a []int, b []int) []int {
	n := len(a) + len(b)
	merged := make([]int, n)

	for f, l, r := 0, 0, 0; l+r < n; f++ {
		if l == len(a) {
			copy(merged[f:], b[r:])
			break
		}
		if r == len(b) {
			copy(merged[f:], a[l:])
			break
		}
		if a[l] < b[r] {
			merged[f] = a[l]
			l++
		} else {
			merged[f] = b[r]
			r++
		}
	}
	return merged
}
