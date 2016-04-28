package main

// merge uses 2 finger algorithms to merge 2 sorted arrays.
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
