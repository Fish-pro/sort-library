/*
Copyright 2021
author: York

归并排序
原理：将数组拆分，然后排序后再合并
时间复杂度：O(n)
空间复杂度：O(log2^n)
*/

package sort_go

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := MergeSort(nums[0:i])
	right := MergeSort(nums[i:])

	return merge(left, right)
}

func merge(l, r []int) []int {
	i, j := 0, 0
	m, n := len(l), len(r)
	var res []int
	for i < m && j < n {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, r[j:]...)

	return res
}
