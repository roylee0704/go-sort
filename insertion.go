package main

// insertion costs O(n * j) of time complexity
func insertion(datas []int) {

	// O(n) steps
	for key := 1; key < len(datas); key++ {

		// O(j) compares & swaps
		pairwiseSwap(datas, key)
	}
}

// pairwiseSwap in ascending order
func pairwiseSwap(datas []int, key int) {

	var temp int
	for j := len(datas) - 1; j >= 0 && key >= 0; j-- {
		temp = datas[j]

		// O(j) compare
		if datas[key] > temp {
			continue
		}
		// O(j) swap
		datas[j] = datas[key]
		datas[key] = temp
		key = j
	}

}
