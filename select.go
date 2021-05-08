/*
Copyright 2021
author: York

选择排序
原理：
对于给定的一组纪录，经过比较后得到最小的记录，然后将该记录与第一个记录镜像交换，
然后将除了第一个记录的其他值进行同样的操作
时间复杂度：O(n^2)
空间复杂度：O(1)
*/

package sort_go

func SelectSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		tmp := array[i]
		flag := i
		for j := i + 1; j < len(array); j++ {
			if tmp > array[j] {
				tmp = array[j]
				flag = j
			}
		}
		if flag != i {
			array[flag] = array[i]
			array[i] = tmp
		}
	}
	return array
}
