/*
Copyright 2021
author: York

插入排序
原理：从已有的记录中不断取出元素，插入到新的有序列表中
时间复杂度：O(n^2)
空间复杂度：O(1)
*/

package sort_go

func InsertSort(array []int) []int {
	if len(array) == 0 {
		return array
	}
	for i := 1; i < len(array); i++ {
		tmp, j := array[i], i
		// 如果元素比排序后最大值大，那么就得移动元素
		if array[j-1] > tmp {
			for j >= 1 && array[j-1] > tmp {
				array[j] = array[j-1]
				j--
			}
			array[j] = tmp
		}
	}
	return array
}
