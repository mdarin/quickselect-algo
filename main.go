package main

import (
	"fmt"
	"math/rand/v2"
)

// Input: [7, 4, 6, 3, 9, 1]
// k = 2
// Output: k’th smallest array element is 3
//
// Input: [7, 4, 6, 3, 9, 1]
// k = 1
// Output: k’th smallest array element is 1

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// Please note that this returns numbers in the range [min, max), meaning that max is never returned.
func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

// Partition using Lomuto partition scheme.
func partition(a []int, left int, right int, pIndex int) int {
	// pick `pIndex` as a pivot from the array
	pivot := a[pIndex]

	// Move pivot to end
	swap(a, pIndex, right)

	// elements less than the pivot will be pushed to the left of `pIndex`;
	// elements more than the pivot will be pushed to the right of `pIndex`;
	// equal elements can go either way
	pIndex = left

	// each time we find an element less than or equal to the pivot, `pIndex`
	// is incremented, and that element would be placed before the pivot.
	for i := left; i < right; i++ {
		if a[i] <= pivot {
			swap(a, i, pIndex)
			pIndex++
		}
	}

	// move pivot to its final place
	swap(a, pIndex, right)

	// return `pIndex` (index of the pivot element)
	return pIndex
}

// Returns the k'th smallest element in the list within `left…right` (inclusive).
func QuickSelect(nums []int, k int) int {
	left := 1
	right := len(nums) - 1

	for {
		// If the array contains only one element, return that element
		if left == right {
			return nums[left]
		}

		// select `pIndex` between left and right
		pIndex := randRange(left, right)

		pIndex = partition(nums, left, right, pIndex)

		// The pivot is in its final sorted position
		if k == pIndex {
			return nums[k]
		}

		// if `k` is less than the pivot index
		if k < pIndex {
			right = pIndex - 1

			continue
		}

		// if `k` is more than the pivot index
		left = pIndex + 1
	}
}

func main() {
	nums := []int{7, 4, 6, 3, 9, 1}
	selected := QuickSelect(nums, 2)

	fmt.Println("selected", selected)
}
