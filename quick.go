/*
Copyright 2021
author: York

快速排序
原理：
时间复杂度：
空间复杂度：
*/

package sort_go

func quickSort(values []int, left int, right int) {
	key := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= key {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= key && i <= p {
			i++
		}
		if i < p {
			values[p] = values[i]
			p = i
		}
		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}
}

func QuickSort(values []int) []int {
	quickSort(values, 0, len(values)-1)
	return values
}
