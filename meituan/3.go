package main

import "sort"

func tour(m int, chocolate []int) int {
	n := len(chocolate)
	sort.Ints(chocolate)
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + chocolate[i-1]*chocolate[i-1]
	}
	// pos := sort.Search(len(pre), func(i int) bool {
	// 	return pre[i] > m
	// })
	pos := findLower(pre, m)
	return pos
	// return pos - 1
}

func findLower(nums []int, target int) int {
	//find the last index of nums[i] <= target
	// equals to find the first index of nums[i] > target - 1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// r is the first index of nums[i] > target - 1, l is the first index of nums[i] > target
	return r
}
