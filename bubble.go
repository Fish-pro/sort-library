/*
Copyright 2021
author: York

冒泡排序
原理：地带n次，从后往前，如果后一个值大雨前一个值，则交换位置
时间复杂度：O(n^2)
空间复杂度：O(1)
*/

package sort_go

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := len(array) - 1; j > i; j-- {
			if array[j] < array[j-1] {
				tmp := array[j]
				array[j] = array[j-1]
				array[j-1] = tmp
			}
		}
	}
	return array
}
